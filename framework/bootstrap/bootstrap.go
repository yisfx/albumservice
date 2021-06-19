package bootstrap

import (
	"albumservice/framework/constFiled"
	"albumservice/framework/model"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strings"

	log "github.com/skoo87/log4go"
)

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

var controllerMap = map[string]map[string]model.RouterMap{}

func Response404(resp http.ResponseWriter, httpMethodName string, request *http.Request) {
	fmt.Println("404 :", httpMethodName, request.URL.Path)
	resss, _ := json.Marshal(fmt.Sprint("404:     ", httpMethodName, ":", request.URL.Path))
	resp.Write(resss)
	return
}

func getRoute(resp http.ResponseWriter, request *http.Request) (*model.RouterMap, bool) {
	isPost := false
	httpMethodName := constFiled.Get
	if strings.EqualFold(request.Method, constFiled.Post) {
		isPost = true
		httpMethodName = constFiled.Post
	}
	urls := strings.Split(request.URL.Path, "/")
	routeBase := urls[2]

	controller, hasController := controllerMap[routeBase]
	if !hasController {
		Response404(resp, httpMethodName, request)
		return nil, false
	}

	route, hasRoute := controller[urls[3]]

	if !hasRoute {
		Response404(resp, httpMethodName, request)
		return nil, false
	}

	if isPost != route.IsPost {
		Response404(resp, httpMethodName, request)
		return nil, false
	}
	return &route, true
}

func Process(resp http.ResponseWriter, request *http.Request) {

	defer func() {
		err := recover()

		switch err.(type) {
		case runtime.Error:
			{ // 运行时错误
				log.Error("err %s", err)
			}
		default:
			{ // 非运行时错误
				log.Error("err %s", err)
			}
		}
	}()

	route, exist := getRoute(resp, request)

	if !exist {
		return
	}

	var result []reflect.Value
	if route.ArgType == nil {
		result = route.Controller.Call(nil)
	} else {
		a := reflect.New(route.ArgType).Interface()
		MustJSONDecode(ReadBody(request.Body), a)
		args := []reflect.Value{reflect.ValueOf(a)}
		result = route.Controller.Call(args)
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
	for _, curController := range ControllerList {

		controller := curController.Controller

		controllerName := curController.ControllerName

		routerList := map[string]model.RouterMap{}
		controllerValue := reflect.ValueOf(controller).Elem().Elem()
		controllerType := reflect.TypeOf(*controller)
		for methodIndex := 0; methodIndex < controllerValue.NumMethod(); methodIndex++ {

			methodType := controllerType.Method(methodIndex)
			methodValue := controllerValue.Method(methodIndex)

			route := &model.RouterMap{}
			route.Controller = methodValue
			post, routeName := isPost(methodType.Name)

			if methodValue.Type().NumIn() > 0 {
				route.ArgType = methodValue.Type().In(0).Elem()
			} else {
				route.ArgType = nil
			}
			route.IsPost = post
			routerList[routeName] = *route
			httpMethod := constFiled.Get
			if post {
				httpMethod = constFiled.Post
			}

			fmt.Println(httpMethod, "router:", routeName, controllerName+routeName, methodValue)
		}
		v := controllerValue.Elem()
		routeField := v.FieldByName("RouterList")
		routeField.Set(reflect.ValueOf(routerList))
		controllerMap[controllerName] = routerList
	}
	http.HandleFunc("/api/", Process)
}
