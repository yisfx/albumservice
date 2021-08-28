package bootstrap

import (
	"albumservice/albumtool"
	"albumservice/framework/constFiled"
	"albumservice/framework/model"
	"albumservice/framework/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

var SysConfig, GlobalConf, albumHelper reflect.Value

func SetConfig(SysConfigModel *model.SysConf, GlobalConfModel *model.GlobalConf) {
	SysConfig = reflect.ValueOf(SysConfigModel)
	GlobalConf = reflect.ValueOf(GlobalConfModel)
	albumHelper = reflect.ValueOf(&albumtool.AlbumHelper{})
}

func MustJSONDecode(b []byte, i interface{}) {
	err := json.Unmarshal(b, i)
	if err != nil {
		panic(err)
	}
}

func isPost(name string) (bool, string) {
	if l := strings.Split(name, "_"); len(l) == 2 {
		return strings.EqualFold(l[0], constFiled.Post), l[1]
	}
	return false, name
}

var ControllerRouterMap = map[string]*ControllerRouteType{}

func Response404(resp http.ResponseWriter, httpMethodName string, request *http.Request) {
	fmt.Println("404 :", httpMethodName, request.URL.Path)
	resss, _ := json.Marshal(fmt.Sprint("404:     ", httpMethodName, ":", request.URL.Path))
	resp.Write(resss)
	return
}

func getRoute(resp http.ResponseWriter, request *http.Request) (*reflect.Value, *RouterCell, bool) {
	defer utils.HanderError("getRoute")
	isPost := false
	httpMethodName := constFiled.Get
	if strings.EqualFold(request.Method, constFiled.Post) {
		isPost = true
		httpMethodName = constFiled.Post
	}
	urls := strings.Split(request.URL.Path, "/")
	routeBase := urls[2]

	controller, hasController := ControllerRouterMap[routeBase]
	if !hasController {
		Response404(resp, httpMethodName, request)
		return nil, nil, false
	}

	routeCell, hasRoute := controller.RouteFunc[urls[3]]

	if !hasRoute {
		Response404(resp, httpMethodName, request)
		return nil, nil, false
	}

	if isPost != routeCell.IsPost {
		Response404(resp, httpMethodName, request)
		return nil, nil, false
	}

	//new controller
	controllerVale := reflect.New(controller.ControllerType.Elem())
	controllerVale.Elem().FieldByName("SysConfig").Set(SysConfig)
	controllerVale.Elem().FieldByName("GlobalConf").Set(GlobalConf)
	controllerVale.Elem().FieldByName("AlbumHelper").Set(albumHelper)

	///set controller field
	// SysConfig  model.SysConf
	// GlobalConf model.GlobalConf
	// albumHelper *albumtool.AlbumHelper

	routeMethod := controllerVale.MethodByName(httpMethodName + "_" + urls[3])

	return &routeMethod, routeCell, true
}

func Process(resp http.ResponseWriter, request *http.Request) {

	defer utils.HanderError("Process")

	routerMethod, routerCell, exist := getRoute(resp, request)

	if !exist {
		return
	}

	var result []reflect.Value
	if routerCell.ArgType == nil {
		result = routerMethod.Call(nil)
	} else {
		a := reflect.New(routerCell.ArgType).Interface()
		MustJSONDecode(ReadBody(request.Body), a)
		args := []reflect.Value{reflect.ValueOf(a)}
		result = routerMethod.Call(args)
	}
	if result == nil {
		resp.Write(([]byte)("are you ok?"))
	} else {
		r, err := json.Marshal(result[0].Interface())
		if err != nil {
			fmt.Println("err:", err)
		}
		resp.Write(r)
	}
}

func Bootstrap(ControllerList ...model.ControllerData) {
	defer utils.ErrorHandler()
	for _, curController := range ControllerList {

		controllerName := curController.ControllerName

		routerList := &ControllerRouteType{}
		routerList.RouteFunc = map[string]*RouterCell{}
		routerList.ControllerType = curController.ControllerType

		var controllerValue reflect.Value = reflect.New(curController.ControllerType).Elem()
		controllerType := curController.ControllerType
		fmt.Println(controllerName, " methods:", controllerValue.NumMethod())
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
	http.HandleFunc("/api/", Process)
}
