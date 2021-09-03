package bootstrapmodel

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Request      *http.Request
	RequestBody  interface{}
	ResponseBody interface{}
	HttpCode     int
	response     http.ResponseWriter

	hasSendResponse bool
}

func NewContext(request *http.Request, response *http.ResponseWriter) *Context {
	return &Context{Request: request, response: *response}
}

func (context *Context) getHeader(key string) string {
	if v, ok := context.Request.Header[key]; ok {
		if len(v) > 0 {
			return v[0]
		}
	}
	return ""
}

func (context *Context) getQueryString(key string) string {
	if v, ok := context.Request.URL.Query()[key]; ok {
		if len(v) > 0 {
			return v[0]
		}
	}
	return ""
}

func (context *Context) GetParam(key string) string {
	v := context.getQueryString(key)
	if v == "" {
		return context.getHeader(key)
	}
	return v
}

func (context *Context) ResponseSend() {
	if context.hasSendResponse {
		return
	}
	if context.ResponseBody == nil {
		context.ResponseBody = "hello body~!"
	}
	r, err := json.Marshal(context.ResponseBody)
	if err != nil {
		fmt.Println("err:", err)
	}
	if context.HttpCode != 0 {
		context.response.WriteHeader(context.HttpCode)
	}
	context.hasSendResponse = true
	context.response.Write(r)
}
