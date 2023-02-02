package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.gf.com.cn/hk-common/go-boot/dto/base"
	"gitlab.gf.com.cn/hk-common/go-boot/isp"
	"net/http"
)

// MCheckLogin 返回登录中间件
func MCheckLogin(serverKey, loginAPIPublic, userAPI string) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		tokenC, _ := c.Request.Cookie("access_token")
		token := ""
		if tokenC != nil {
			token = tokenC.Value
		} else {
			token = c.GetHeader("access_token")
		}
		if token == "" {
			c.Abort()
			r := base.Result{
				ErrCode: 0,
				ErrMsg:  "",
				Data:    "请登录！",
			}
			c.JSON(http.StatusUnauthorized, r)
			return
		}

		acl, err := isp.GetUserinfo(token, serverKey, loginAPIPublic, userAPI)
		if err != nil {
			isp.ClearCookie(c.Writer)
			c.Abort()
			r := base.Result{
				ErrCode: 0,
				ErrMsg:  "",
				Data:    err.Error(),
			}
			c.JSON(http.StatusUnauthorized, r)
			return
		}

		if acl.ID == "" {
			c.Abort()
			r := base.Result{
				ErrCode: 0,
				ErrMsg:  "",
				Data:    "未登录或token已过期！",
			}
			c.JSON(http.StatusUnauthorized, r)
			return
		}

		loginInfo := isp.LoginInfo{
			Birthday:       acl.Birthday,
			BranchNo:       acl.BranchNo,
			BusinessDuty:   acl.BusinessDuty,
			BusinessGroup:  acl.BusinessGroup,
			CancelDate:     acl.CancelDate,
			ChineseName:    acl.ChineseName,
			CompanyNo:      acl.CompanyNo,
			DepNo:          acl.DepNo,
			Email:          acl.Email,
			EnglishName:    acl.EnglishName,
			ErpNo:          acl.ErpNo,
			HomePhone:      acl.HomePhone,
			ID:             acl.ID,
			IDdNo:          acl.IDdNo,
			IDType:         acl.IDType,
			LoginID:        acl.LoginID,
			MobilePhone:    acl.MobilePhone,
			OfficePhone:    acl.OfficePhone,
			OpenDate:       acl.OpenDate,
			PasswordUpdate: acl.PasswordUpdate,
			Sex:            acl.Sex,
			Status:         acl.Status,
			UserType:       acl.UserType,
			Nodes:          acl.Nodes,
			Process:        acl.ProcessQueries,
			ProcessOps:     acl.ProcessOps,
			Menus:          acl.Menus,
			MenuOps:        acl.MenuOps,
		}
		c.Set("user", loginInfo)
		c.Next()
	}
}
