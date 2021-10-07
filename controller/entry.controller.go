package controller

import (
	"albumservice/albumtool/albumUtils"
	"albumservice/framework/bootstrap"
	"albumservice/framework/fxfilter"
	"albumservice/model/request"
	"albumservice/model/response"
)

type EntryController struct{}

func NewEntryController() bootstrap.BaseController {
	return &EntryController{}
}
func (controller EntryController) GetFilterMapping() fxfilter.FilterMapping {
	mapping := fxfilter.FilterMapping{}

	mapping["DeEntry"] = fxfilter.FilterFuncList{}

	return mapping
}

func (controller *EntryController) Post_DeEntry(r *request.DeEntryImageRequest) *response.DeEntryImageResponse {

	result := response.DeEntryImageResponse{}
	i := albumUtils.DecryptImageUri(r.V)
	result.Result = true
	result.Image = *i
	return &result
}
