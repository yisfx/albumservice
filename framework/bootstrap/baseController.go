package bootstrap

import (
	"reflect"
)

type BaseController interface {
	GetFilterMapping() FilterMapping
}

type ControllerData struct {
	ControllerName string
	ControllerType reflect.Type
	///route with filter
	FilterMapper FilterMapping
}

func NewControllerData(ControllerName string, Controller BaseController) *ControllerData {
	o := &ControllerData{}
	o.ControllerName = ControllerName
	o.ControllerType = reflect.TypeOf(Controller)
	o.FilterMapper = Controller.GetFilterMapping()

	return o
}
