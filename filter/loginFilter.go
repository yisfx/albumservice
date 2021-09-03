package filter

import (
	"albumservice/albumtool/loginHelper"
	"albumservice/framework/bootstrapmodel"
	"fmt"
)

func LoginFilter(context *bootstrapmodel.Context) bool {
	fmt.Println("fx-login-token:", context.Request.Header.Get("fx-login-token"))

	loginToken := context.Request.Header.Get("fx-login-token")

	return loginHelper.ValidateLoginStatus(loginToken)
}
