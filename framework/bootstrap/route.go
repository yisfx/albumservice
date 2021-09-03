package bootstrap

import (
	"albumservice/framework/constFiled"
	"albumservice/framework/fxfilter"
	"reflect"
)

type ControllerRouteType struct {
	RouteFunc      map[string]*RouterCell
	ControllerType reflect.Type
}

type RouterCell struct {
	ArgType    reflect.Type
	IsPost     bool
	FilterList fxfilter.FilterFuncList
}

func (r RouterCell) HttpMethod() string {
	if r.IsPost {
		return constFiled.Post
	}
	return constFiled.Get
}
