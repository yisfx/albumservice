package bootstrap

import (
	"albumservice/albumtool"
	"albumservice/framework/bootstrapmodel"
	"fmt"
)

var controllerFieldMap map[string]interface{} = make(map[string]interface{})

func SetConfig(sysConfig bootstrapmodel.SysConf, globalConfig bootstrapmodel.GlobalConf) {
	field := map[string]interface{}{}
	field["SysConfig"] = sysConfig
	field["GlobalConf"] = globalConfig
	field["AlbumHelper"] = albumtool.AlbumHelper{}
	SetControllerFiled(field)
}

func SetControllerFiled(fieldMap map[string]interface{}) {
	for fn, _ := range fieldMap {
		fv := fieldMap[fn]
		controllerFieldMap[fn] = fv
		fmt.Println("inject field:", fn)
	}
}
