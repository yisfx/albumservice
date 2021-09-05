package controller

import (
	"albumservice/framework/bootstrap"
	"albumservice/framework/bootstrapmodel"
	"albumservice/framework/fxfilter"
	"fmt"
)

type DemoController struct {
	Context *bootstrapmodel.Context
}

func (dc *DemoController) GetFilterMapping() fxfilter.FilterMapping {
	mapping := fxfilter.FilterMapping{}
	// mapping["Demo3"] = fxfilter.FilterFuncList{filter.LoginFilter}
	return mapping
}

func (dc *DemoController) Get_Demo1() {

}
func (dc *DemoController) Get_Demo2() {

}

func (dc *DemoController) Post_Demo3(request *bootstrapmodel.BaseResponse) *bootstrapmodel.BaseResponse {
	fmt.Printf("demo request %#v \n", request)
	fmt.Println(dc.Context.Request.URL.Path)

	dc.Context.ResponseBody = "ffffffffffff"
	dc.Context.HttpCode = 408
	dc.Context.ResponseSend()
	return &bootstrapmodel.BaseResponse{Result: true}
}

func NewDemoController() bootstrap.BaseController {
	o := &DemoController{}
	return o
}
