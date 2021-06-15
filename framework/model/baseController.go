package model

import (
	"reflect"
)

type RouterList [10]RouterMap

type RouterMap struct {
	ArgType    reflect.Type
	Controller reflect.Value
	IsPost     bool
}

type BaseController interface {
}

type ControllerData struct {
	ControllerName string
	Controller     *BaseController
}

func NewControllerData(ControllerName string, Controller *BaseController) *ControllerData {
	o := &ControllerData{}
	o.ControllerName = ControllerName
	o.Controller = Controller
	return o
}
