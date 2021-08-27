package bootstrap

import "reflect"

type ControllerRouteType struct {
	RouteFunc      map[string]*RouterCell
	ControllerType reflect.Type
}

type RouterCell struct {
	ArgType reflect.Type
	IsPost  bool
}
