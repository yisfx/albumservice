package controller

import (
	"albumservice/albumtool/constfield"
	"albumservice/albumtool/loginHelper"
	"albumservice/framework/bootstrap"
	"albumservice/framework/bootstrapmodel"
	"albumservice/framework/fxfilter"
	"albumservice/model/request"
	"albumservice/model/response"
	"strings"
)

type LoginController struct {
	GlobalConf bootstrapmodel.GlobalConf
	Context    *bootstrapmodel.Context
}

func NewLoginController() bootstrap.BaseController {
	o := &LoginController{}
	return o
}

func (controller *LoginController) GetFilterMapping() fxfilter.FilterMapping {

	mapping := fxfilter.FilterMapping{}
	// mapping["Login"] = fxfilter.FilterFuncList{filter.LoginFilter}
	return mapping
}

func (controller *LoginController) Post_Login(r *request.LoginRequest) *response.LoginResponse {
	result := &response.LoginResponse{}
	for k, v := range controller.GlobalConf.AdminPwd {
		if p, ok := r.Password[k]; !ok || !strings.EqualFold(p, v) {
			result.Result = false
			result.ErrorMessage = "password error"
			return result
		}
	}
	result.Result = true
	result.LoginToken = loginHelper.SaveLoginToken(r.Password, r.IP)
	return result
}

func (controller *LoginController) Post_Auth() *bootstrapmodel.BaseResponse {
	result := &bootstrapmodel.BaseResponse{}
	result.Result = loginHelper.ValidateLoginStatus(controller.Context.GetParam(constfield.Header_Login_Token_Key))
	if !result.Result {
		result.ErrorMessage = "access invalid"
	}
	return result
}
