package interceptor

import (
	"fmt"
	"net/http"
	"reflect"
)

type LoginInterceptor struct {
}

func NewLoginInterceptor() LoginInterceptor {
	return LoginInterceptor{}
}

func (lp LoginInterceptor) Interfaceptor(request *http.Request, controller *reflect.Value) bool {
	fmt.Println("fx-login-token:", request.Header.Get("fx-login-token"))
	return true
}
