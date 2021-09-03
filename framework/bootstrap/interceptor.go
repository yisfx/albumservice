package bootstrap

import (
	"albumservice/framework/bootstrapmodel"
	"reflect"
)

var interceptorList []*FXInterceptor = []*FXInterceptor{}

type FXInterceptor interface {
	Interfaceptor(context *bootstrapmodel.Context, controller *reflect.Value) bool
}

func AddInterceptor(interceptor ...FXInterceptor) {
	for _, i := range interceptor {
		interceptorList = append(interceptorList, &i)
	}
}
