package bootstrap

import (
	"albumservice/framework/constFiled"
	"albumservice/framework/utils"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func isPost(name string) (bool, string, bool) {
	if l := strings.Split(name, "_"); len(l) == 2 {
		return strings.EqualFold(l[0], constFiled.Post), l[1], true
	}
	return false, name, false
}

var ControllerRouterMap = map[string]*ControllerRouteType{}

func Bootstrap(ControllerList ...ControllerData) {
	defer utils.HanderError()
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

			post, routeName, isRoute := isPost(methodType.Name)

			if !isRoute {
				continue
			}

			if methodValue.Type().NumIn() > 0 {
				route.ArgType = methodValue.Type().In(0).Elem()
			} else {
				route.ArgType = nil
			}
			route.IsPost = post

			route.FilterList = append(route.FilterList, curController.FilterMapper[routeName]...)
			route.RouterMethodName = methodType.Name
			routerList.RouteFunc[strings.ToLower(routeName)] = route
			httpMethod := constFiled.Get
			if post {
				httpMethod = constFiled.Post
			}
			fmt.Println(httpMethod, "router:", routeName, controllerName+"/"+routeName, methodValue)
		}

		ControllerRouterMap[strings.ToLower(controllerName)] = routerList
	}
	fmt.Println("***************************************************")
	http.HandleFunc("/api/", Process)
}
