package filter

import (
	"albumservice/albumtool/loginHelper"
	"fmt"
	"net/http"
)

func LoginFilter(request *http.Request) bool {
	fmt.Println("fx-login-token:", request.Header.Get("fx-login-token"))

	loginToken := request.Header.Get("fx-login-token")

	return loginHelper.ValidateLoginStatus(loginToken)
}
