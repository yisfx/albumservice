package bootstrap

import (
	"albumservice/framework/bootstrapmodel"
	"albumservice/framework/constFiled"
	"albumservice/framework/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func ReadBody(body io.Reader) []byte {
	b, _ := ioutil.ReadAll(body)
	return b
}

func getRoute(resp http.ResponseWriter, request *http.Request) (curController *reflect.Value, routerMethod *reflect.Value, routeCell *RouterCell, exists bool) {
	isPost := false
	httpMethodName := constFiled.Get
	if strings.EqualFold(request.Method, constFiled.Post) {
		isPost = true
		httpMethodName = constFiled.Post
	}
	urls := strings.Split(request.URL.Path, "/")
	routeBase := urls[2]

	controller, hasController := ControllerRouterMap[strings.ToLower(routeBase)]
	if !hasController {
		return nil, nil, nil, false
	}

	routeCell, hasRoute := controller.RouteFunc[strings.ToLower(urls[3])]

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

func preProcess(context *bootstrapmodel.Context, controllerValue *reflect.Value, routerCell *RouterCell) bool {
	/// interceptor
	for _, inter := range interceptorList {
		if !(*inter).Interfaceptor(context, controllerValue) {
			Response401(context)
			return false
		}
	}

	///filter
	for _, filter := range routerCell.FilterList {
		if !filter(context) {
			Response401(context)
			return false
		}
	}

	/// inject controller fields
	InjectControllerField(controllerValue, context, controllerFieldMap)
	return true
}

func Process(resp http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := utils.HanderError(); err != nil {
			fmt.Println(500)
			resp.WriteHeader(500)
			resp.Write([]byte("service error!"))
		}
	}()

	controllerValue, routerMethod, routerCell, exist := getRoute(resp, request)

	context := bootstrapmodel.NewContext(request, &resp)

	if !exist {
		Response404(context)
		return
	}

	/// preProcess
	if !preProcess(context, controllerValue, routerCell) {
		return
	}

	var args []reflect.Value = nil
	if routerCell.ArgType != nil {
		a := reflect.New(routerCell.ArgType).Interface()
		json.Unmarshal(ReadBody(request.Body), a)
		context.RequestBody = a
		args = []reflect.Value{reflect.ValueOf(context.RequestBody)}
	}

	/// exec controller
	result := routerMethod.Call(args)

	if result != nil {
		context.ResponseBody = result[0].Interface()
	}
	///write response
	context.ResponseSend()
}
