package filter

import (
	"albumservice/albumtool/loginHelper"
	"albumservice/framework/bootstrapmodel"
)

func LoginFilter(context *bootstrapmodel.Context) bool {

	loginToken := context.Request.Header.Get("fx-login-token")

	return loginHelper.ValidateLoginStatus(loginToken)
}
