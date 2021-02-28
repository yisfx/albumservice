package modal

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
