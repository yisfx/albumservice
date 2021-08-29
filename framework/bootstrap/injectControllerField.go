package bootstrap

import (
	"albumservice/framework/utils"
	"net/http"
	"reflect"
)

func InjectControllerField(controllerVale *reflect.Value, request *http.Request, fields map[string]interface{}) *reflect.Value {
	defer utils.HanderError("InjectControllerField")

	for fieldName, _ := range fields {
		fv := controllerVale.Elem().FieldByName(fieldName)
		if fv.CanSet() {
			fv.Set(reflect.ValueOf(fields[fieldName]))
		}
	}

	fv := controllerVale.Elem().FieldByName("Request")
	if fv.CanSet() {
		fv.Set(reflect.ValueOf(request))
	}
	return controllerVale
}
