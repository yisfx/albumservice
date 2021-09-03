package bootstrap

import (
	"albumservice/framework/bootstrapmodel"
	"reflect"
)

func InjectControllerField(controllerVale *reflect.Value, context *bootstrapmodel.Context, fields map[string]interface{}) *reflect.Value {

	for fieldName, _ := range fields {
		fv := controllerVale.Elem().FieldByName(fieldName)
		if fv.CanSet() {
			fv.Set(reflect.ValueOf(fields[fieldName]))
		}
	}

	fv := controllerVale.Elem().FieldByName("Context")
	if fv.CanSet() {
		fv.Set(reflect.ValueOf(context))
	}
	return controllerVale
}
