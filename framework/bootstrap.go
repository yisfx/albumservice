package framework

import (
	model "albumservice/model"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func MustJSONDecode(b []byte, i interface{}) {
	err := json.Unmarshal(b, i)
	if err != nil {
		panic(err)
	}
}

func isPost(name string) (bool, string) {
	if l := strings.Split(name, "_"); len(l) == 2 {
		return strings.EqualFold(l[0], "Post"), l[1]
	}
	return false, name
}

func Bootstrap(o *model.BaseController, routeBase string, handler func(http.ResponseWriter, *http.Request)) {

	routerList := map[string]model.RouterMap{}
	controllerValue := reflect.ValueOf(o).Elem().Elem()
	controllerType := reflect.TypeOf(*o)
	for methodIndex := 0; methodIndex < controllerValue.NumMethod(); methodIndex++ {

		methodType := controllerType.Method(methodIndex)
		methodValue := controllerValue.Method(methodIndex)

		if strings.EqualFold(methodType.Name, "process") {
			continue
		}

		route := &model.RouterMap{}
		route.Controller = methodValue
		post, routeName := isPost(methodType.Name)

		if methodValue.Type().NumIn() > 0 {
			route.ArgType = methodValue.Type().In(0).Elem()
		} else {
			route.ArgType = nil
		}
		routerList[routeName] = *route

		fmt.Println(post, "router:", routeName, routeBase+routeName, methodValue)
	}
	v := controllerValue.Elem()
	routeField := v.FieldByName("RouterList")
	routeField.Set(reflect.ValueOf(routerList))
	http.HandleFunc(routeBase, handler)
}
