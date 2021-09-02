package bootstrap

import (
	"net/http"
)

type FilterMapping map[string]FilterFuncList

type FilterFuncList []FilterFunc

type FilterFunc func(request *http.Request) bool

type FXFilter interface {
	Filter(request *http.Request) bool
}
