package controller

import (
	"albumservice/albumtool/loginHelper"
	"albumservice/framework/bootstrap"
	"albumservice/framework/model"
	"albumservice/model/request"
	"albumservice/model/response"
	"net/http"
	"strings"
)

type LoginController struct {
	GlobalConf model.GlobalConf
	Request    *http.Request
}

func NewLoginController() bootstrap.BaseController {
	o := &LoginController{}
	return o
}

func (controller *LoginController) GetFilterMapping() bootstrap.FilterMapping {

	mapping := bootstrap.FilterMapping{}

	return mapping
}

func (controller *LoginController) Post_Login(r *request.LoginRequest) (result *response.LoginResponse) {

	for k, v := range controller.GlobalConf.AdminPwd {
		if p, ok := r.Password[k]; !ok || !strings.EqualFold(p, v) {
			result.Result = false
			result.ErrorMessage = "password error"
			return result
		}
	}

	result.LoginToken = loginHelper.SaveLoginToken(r.Password, r.IP)
	return result
}
