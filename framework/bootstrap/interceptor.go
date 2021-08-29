package bootstrap

import (
	"net/http"
	"reflect"
)

var interceptorList []*FXInterceptor = []*FXInterceptor{}

type FXInterceptor interface {
	Interfaceptor(request *http.Request, controller *reflect.Value) bool
}

func AddInterceptor(interceptor ...*FXInterceptor) {
	interceptorList = append(interceptorList, interceptor...)
}
