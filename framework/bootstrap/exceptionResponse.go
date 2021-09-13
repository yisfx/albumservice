package bootstrap

import (
	"albumservice/framework/bootstrapmodel"
	"fmt"
)

func Response404(context *bootstrapmodel.Context) {
	resss := fmt.Sprint("404:     ", context.Request.Method, ":", context.Request.URL.Path)
	context.HttpCode = 404
	context.ResponseBody = resss
	context.ResponseSend()
}

func Response401(context *bootstrapmodel.Context) {
	resss := fmt.Sprint("401:", context.Request.Method, ":", context.Request.URL.Path)
	context.ResponseBody = resss
	context.HttpCode = 401
	context.ResponseSend()
}

func Response500(context *bootstrapmodel.Context) {
	resss := fmt.Sprint("500:", context.Request.Method, ":", context.Request.URL.Path)
	context.ResponseBody = resss
	context.HttpCode = 500
	context.ResponseSend()
}
