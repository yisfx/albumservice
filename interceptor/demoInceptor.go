package interceptor

import (
	"net/http"
	"reflect"
)

type DemoInterceptor struct {
}

func NewDemoInterceptor() DemoInterceptor {
	return DemoInterceptor{}
}

func (dp DemoInterceptor) Interfaceptor(request *http.Request, controller *reflect.Value) bool {
	return true
}
