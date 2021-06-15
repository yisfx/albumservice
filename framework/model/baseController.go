package model

import (
	"net/http"
	"reflect"
)

type RouterList [10]RouterMap

type RouterMap struct {
	ArgType    reflect.Type
	Controller reflect.Value
	IsPost     bool
}

type BaseController interface {
	Process(res http.ResponseWriter, request *http.Request)
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
