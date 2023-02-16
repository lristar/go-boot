package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	myerror "gitlab.gf.com.cn/hk-common/go-boot/lib/error"
	"gitlab.gf.com.cn/hk-common/go-boot/logger"
	"gitlab.gf.com.cn/hk-common/go-boot/pkg/sentry"
	"io"
	"net/http"
	//logger "platform-backend/server/log"
	"strings"

	"github.com/gin-gonic/gin"
)

// PrintBody 返回打印body中间件
func PrintBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.ErrorStack()
				sentry.LogAndSentry(fmt.Errorf("%v", err))
				if mErr, ok := err.(myerror.BootError); ok {
					logger.ErrorHandler(c, mErr.Code(), mErr)
				} else {
					logger.ErrorHandler(c, http.StatusNotAcceptable, fmt.Errorf("%v", err))
				}
			}
		}()

		body := ""
		if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") && c.Request.RequestURI != "/login" {
			byteBody, _ := io.ReadAll(c.Request.Body)
			if len(byteBody) != 0 {
				// 去掉body传参中的字符串前后空格（只处理第一层）
				mapBody := make(map[string]interface{})
				if err := json.Unmarshal(byteBody, &mapBody); err == nil {
					for key, value := range mapBody {
						if v, ok := value.(string); ok {
							mapBody[key] = strings.TrimSpace(v)
						}
					}
					byteBody, _ = json.Marshal(mapBody)
				}
				// 重新给body赋值
				c.Request.Body = io.NopCloser(bytes.NewBuffer(byteBody))
				body = "" + strings.Replace(strings.Replace(string(byteBody), "\n", "", -1), "\t", "", -1)
			}
		}

		logger.GinLog(c.Request.Method, c.Request.RequestURI, body)
		c.Next()
	}
}
