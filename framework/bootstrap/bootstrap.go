package bootstrap

import (
	"albumservice/albumtool"
	"albumservice/framework/constFiled"
	"albumservice/framework/model"
	"albumservice/framework/utils"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

var controllerFieldMap map[string]interface{} = make(map[string]interface{})

func SetConfig(sysConfig model.SysConf, globalConfig model.GlobalConf) {
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

func isPost(name string) (bool, string) {
	if l := strings.Split(name, "_"); len(l) == 2 {
		return strings.EqualFold(l[0], constFiled.Post), l[1]
	}
	return false, name
}

var ControllerRouterMap = map[string]*ControllerRouteType{}

func Bootstrap(ControllerList ...model.ControllerData) {
	defer utils.ErrorHandler()
	fmt.Println("***************************************************")
	for _, curController := range ControllerList {

		controllerName := curController.ControllerName

		routerList := &ControllerRouteType{}
		routerList.RouteFunc = map[string]*RouterCell{}
		routerList.ControllerType = curController.ControllerType

		var controllerValue reflect.Value = reflect.New(curController.ControllerType).Elem()
		controllerType := curController.ControllerType
		fmt.Println("controller :", controllerName, ", methods:", controllerValue.NumMethod())
		for methodIndex := 0; methodIndex < controllerValue.NumMethod(); methodIndex++ {

			route := &RouterCell{}

			methodType := controllerType.Method(methodIndex)
			methodValue := controllerValue.Method(methodIndex)

			post, routeName := isPost(methodType.Name)

			if methodValue.Type().NumIn() > 0 {
				route.ArgType = methodValue.Type().In(0).Elem()
			} else {
				route.ArgType = nil
			}
			route.IsPost = post

			routerList.RouteFunc[routeName] = route
			httpMethod := constFiled.Get
			if post {
				httpMethod = constFiled.Post
			}
			fmt.Println(httpMethod, "router:", routeName, controllerName+"/"+routeName, methodValue)
		}

		ControllerRouterMap[controllerName] = routerList
	}
	fmt.Println("***************************************************")
	http.HandleFunc("/api/", Process)
}
