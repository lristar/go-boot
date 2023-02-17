package lsp

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.gf.com.cn/hk-common/go-boot/lib/http_client"
	"gitlab.gf.com.cn/hk-common/go-boot/pkg/logger"
	"net/http"
)

// LoginInfo 登录用户信息
type LoginInfo struct {
	Birthday       string `json:"birthday"`
	BranchNo       int    `json:"branch_no"`
	BusinessDuty   string `json:"business_duty"`
	BusinessGroup  string `json:"business_group"`
	CancelDate     string `json:"cancel_date"`
	ChineseName    string `json:"chinese_name"`
	CompanyNo      string `json:"company_no"`
	DepNo          string `json:"dep_no"`
	Email          string `json:"email"`
	EnglishName    string `json:"english_name"`
	ErpNo          string `json:"erp_no"`
	HomePhone      string `json:"home_phone"`
	ID             string `json:"id"`
	IDdNo          string `json:"id_no"`
	IDType         string `json:"id_type"`
	LoginID        string `json:"login_id"`
	MobilePhone    string `json:"mobile_phone"`
	OfficePhone    string `json:"office_phone"`
	OpenDate       string `json:"open_date"`
	PasswordUpdate string `json:"password_update"`
	Sex            string `json:"sex"`
	Status         string `json:"status"`
	UserType       string `json:"user_type"`
	Nodes          []string
	Process        []string // 流程查看权限列表
	ProcessOps     []string // 流程操作权限列表
	Menus          []string // 菜单查询权限列表
	MenuOps        []string // 菜单高级权限列表
}

// UserInfo 用户及权限
type UserInfo struct {
	Birthday       string   `json:"birthday"`
	BranchNo       int      `json:"branch_no"`
	BusinessDuty   string   `json:"business_duty"`
	BusinessGroup  string   `json:"business_group"`
	CancelDate     string   `json:"cancel_date"`
	ChineseName    string   `json:"chinese_name"`
	CompanyNo      string   `json:"company_no"`
	DepNo          string   `json:"dep_no"`
	Email          string   `json:"email"`
	EnglishName    string   `json:"english_name"`
	ErpNo          string   `json:"erp_no"`
	HomePhone      string   `json:"home_phone"`
	ID             string   `json:"id"`
	IDdNo          string   `json:"id_no"`
	IDType         string   `json:"id_type"`
	LoginID        string   `json:"login_id"`
	MobilePhone    string   `json:"mobile_phone"`
	OfficePhone    string   `json:"office_phone"`
	OpenDate       string   `json:"open_date"`
	PasswordUpdate string   `json:"password_update"`
	Sex            string   `json:"sex"`
	Status         string   `json:"status"`
	UserType       string   `json:"user_type"`
	Menus          []string `json:"menus"`
	Nodes          []string `json:"nodes"`
	Quotas         []string `json:"quotas"`
	ProcessQueries []string `json:"processQueries"`
	ProcessOps     []string `json:"processOps"`
	MenuOps        []string `json:"menuOps"`
}

// GetUserinfo 获取用户及权限信息
func GetUserinfo(token, serverKey, loginAPIPublic, userAPI string) (*UserInfo, error) {
	request := http_client.NewRequest()
	request.URL = loginAPIPublic + "/userinfo"
	request.Header = map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + token,
	}
	res, err := request.HTTPDo()
	if err != nil {
		logger.Errorf("req.url=%s,req.header.Authorization=%s,res=%s,err=%s", request.URL, request.Header["Authorization"], res, err.Error())
		return nil, err
	}
	info := new(UserInfo)
	err = json.Unmarshal(res, info)
	if err != nil {
		return nil, err
	}
	auths, err := FindUserAuths(userAPI, serverKey, info.ID)
	if err != nil {
		return nil, err
	}
	info.MenuOps, info.Menus = auths.MenuOps, auths.Menus
	info.ProcessOps, info.ProcessQueries, info.Nodes = auths.ProcessOps, auths.Process, auths.Nodes
	return info, nil
}

// FindUserAuths 获取用户权限
func FindUserAuths(userAPI, serverKey, userID string) (Auths, error) {
	request := http_client.NewRequest()
	request.URL = userAPI + "/external/user_auths?user_id=" + userID + "&client_id=" + serverKey
	system, err := request.HTTPDo()
	if err != nil {
		return Auths{}, err
	}
	res := struct {
		Data Auths `json:"data"`
	}{}

	err = json.Unmarshal(system, &res)
	if err != nil {
		return Auths{}, err
	}
	auths := res.Data
	return auths, nil
}

type Auths struct {
	Nodes      []string `json:"nodes"`
	Process    []string `json:"processQueries"` // 流程查看权限列表
	ProcessOps []string `json:"processOps"`     // 流程操作权限列表
	Menus      []string `json:"menus"`          // 菜单查询权限列表
	MenuOps    []string `json:"menuOps"`        // 菜单高级权限列表
}

// ClearCookie 清除cookies
func ClearCookie(writer gin.ResponseWriter) {
	token := &http.Cookie{Name: "access_token", Value: "", Path: "/", MaxAge: -1}
	idToken := &http.Cookie{Name: "id_token", Value: "", Path: "/", MaxAge: -1}
	http.SetCookie(writer, token)
	http.SetCookie(writer, idToken)
}
