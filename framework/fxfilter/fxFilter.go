package fxfilter

import (
	"albumservice/framework/bootstrapmodel"
	"net/http"
)

type FilterMapping map[string]FilterFuncList

type FilterFuncList []FilterFunc

type FilterFunc func(context *bootstrapmodel.Context) bool

type FXFilter interface {
	Filter(request *http.Request) bool
}
