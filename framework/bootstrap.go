package framework

import (
	model "albumservice/model"
	"fmt"
	"reflect"
)

func Bootstrap(o model.BaseController) {
	o.RouterMap = make(map[string]model.RouterMap)
	controller := reflect.TypeOf(0)
	for i := 0; i < controller.NumMethod(); i++ {
		m := controller.Method(i)
		route := model.RouterMap{}
		route.Controller = reflect.ValueOf(m)
		if m.Type.NumIn() > 0 {

			route.ArgType = m.Type.In(1)
		} else {
			route.ArgType = nil
		}
		fmt.Println(route.ArgType)
		o.RouterMap[m.Name] = route
	}
}
