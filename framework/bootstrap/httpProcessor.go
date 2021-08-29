package bootstrap

import (
	"albumservice/framework/constFiled"
	"albumservice/framework/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func Process(resp http.ResponseWriter, request *http.Request) {

	defer utils.HanderError("Process")

	controllerValue, routerMethod, routerCell, exist := getRoute(resp, request)

	if !exist {
		Response404(resp, request.Method, request)
		return
	}
	///inject controller fields

	controllerValue = InjectControllerField(controllerValue, request, controllerFieldMap)
	/// interceptor
	for _, inter := range interceptorList {
		if !(*inter).Interfaceptor(request, controllerValue) {
			Response415(resp, routerCell.HttpMethod(), request)
			return
		}
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

func MustJSONDecode(b []byte, i interface{}) {
	err := json.Unmarshal(b, i)
	if err != nil {
		panic(err)
	}
}

func getRoute(resp http.ResponseWriter, request *http.Request) (curController *reflect.Value, routerMethod *reflect.Value, routeCell *RouterCell, exists bool) {
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
		return nil, nil, nil, false
	}

	routeCell, hasRoute := controller.RouteFunc[urls[3]]

	if !hasRoute {
		return nil, nil, nil, false
	}

	if isPost != routeCell.IsPost {
		return nil, nil, nil, false
	}

	//new controller
	controllerVale := reflect.New(controller.ControllerType.Elem())

	routeMethod := controllerVale.MethodByName(httpMethodName + "_" + urls[3])

	return &controllerVale, &routeMethod, routeCell, true
}
