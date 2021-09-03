package interceptor

import (
	"albumservice/framework/bootstrapmodel"
	"reflect"
)

type DemoInterceptor struct {
}

func NewDemoInterceptor() DemoInterceptor {
	return DemoInterceptor{}
}

func (dp DemoInterceptor) Interfaceptor(context *bootstrapmodel.Context, controller *reflect.Value) bool {
	return true
}
