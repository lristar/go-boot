package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"gitlab.gf.com.cn/hk-common/go-boot/isp"
	"io/ioutil"
	"net/http"
	"strings"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Jaeger() gin.HandlerFunc {
	return func(c *gin.Context) {
		tr := opentracing.GlobalTracer()
		path := c.FullPath()
		for _, param := range c.Params {
			if !strings.Contains(param.Key, "id") {
				path = strings.Replace(path, fmt.Sprintf(":%s", param.Key), param.Value, 1)
			}
		}
		operationName := fmt.Sprintf("%s %s", c.Request.Method, path)
		var parentSpan opentracing.Span
		// 直接从 c.Request.Header 中提取 span,如果没有就新建一个
		spCtx, err := tr.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			parentSpan = tr.StartSpan(operationName)
			defer parentSpan.Finish()
		} else {
			parentSpan = opentracing.StartSpan(
				operationName,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
			defer parentSpan.Finish()
		}
		if c.Request.URL.RawQuery != "" {
			parentSpan.SetTag("http.req.params", c.Request.URL.RawQuery)
		}
		parentSpan.SetTag("http.req.remoteAddr", c.Request.RemoteAddr)
		bts, _ := ioutil.ReadAll(c.Request.Body)
		p := struct {
			ProcessId string `json:"process_id"`
			ClientId  string `json:"client_id"`
		}{}
		json.Unmarshal(bts, &p)
		if p.ProcessId != "" {
			parentSpan.SetTag("process_id", p.ProcessId)
		}
		if p.ClientId != "" {
			parentSpan.SetTag("client_id", p.ClientId)
		}
		parentSpan.LogFields(tracerLog.String("http.req.body", string(bts)))
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bts))
		parentSpan.SetTag("http.req.method", c.Request.Method)
		parentSpan.SetTag("http.req.url", c.Request.RequestURI)
		c.Set("ctx", opentracing.ContextWithSpan(c, parentSpan))
		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		if v, ok := c.Get("user"); ok {
			if info, yes := v.(isp.LoginInfo); yes {
				parentSpan.SetTag("http.user.id", info.ID)
				parentSpan.SetTag("http.user.english_name", info.EnglishName)
				parentSpan.SetTag("http.user.chinese_name", info.ChineseName)
			}

		}
		parentSpan.SetTag("http.res.status_code", c.Writer.Status())
		if c.Writer.Status() != http.StatusOK {
			parentSpan.SetTag("error", true)
			parentSpan.LogFields(tracerLog.String("http.res.body", blw.body.String()))
			bt, _ := json.Marshal(c.Request.Header)
			if bt != nil {
				parentSpan.LogFields(tracerLog.String("http.res.header", string(bt)))
			}
		}
	}
}
