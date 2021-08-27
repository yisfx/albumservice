package model

import (
	"reflect"
)

type BaseController interface {
}

type ControllerData struct {
	ControllerName string
	ControllerType reflect.Type
}

func NewControllerData(ControllerName string, Controller *BaseController) *ControllerData {
	o := &ControllerData{}
	o.ControllerName = ControllerName
	o.ControllerType = reflect.TypeOf(*Controller)
	return o
}
