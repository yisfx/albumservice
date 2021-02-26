package modal

import (
	"reflect"
)

type RouterMap struct {
	ArgType    reflect.Type
	Controller reflect.Value
}

type BaseController struct {
	SysConfig SysConf
	RouterMap map[string]RouterMap
}
