package controller

import (
	"albumservice/model/request"
	"albumservice/model/response"
	"albumservice/framework/fxfilter"
	_ "albumservice/filter"
	"albumservice/framework/bootstrap"
	"albumservice/framework/bootstrapmodel"
	"albumservice/wordtool"
	"albumservice/model"
	"sort"
)

type WordController struct {
	// SysConfig   bootstrapmodel.SysConf
	GlobalConf  bootstrapmodel.GlobalConf
	// AlbumHelper albumtool.AlbumHelper
	Context     *bootstrapmodel.Context
	IsLogin     *bool
}

func NewWordController() bootstrap.BaseController{
	o := &WordController{}
	return o
}


func (controller WordController) GetFilterMapping() fxfilter.FilterMapping {
	mapping := fxfilter.FilterMapping{}

	mapping["GetWord"] = fxfilter.FilterFuncList{}
	mapping["AddWord"] = fxfilter.FilterFuncList{
		// filter.LoginFilter
	}


	return mapping
}

func (this *WordController) GET_GetSection() *response.GetSectionResponse{
	section := wordtool.GetSection(this.GlobalConf.WordFile)

	return &response.GetSectionResponse{Section:section}
}

func (this *WordController) POST_GetWord(request *request.GetWordRequest) *response.GetWordResponse{
	Chapter:= wordtool.GetWord(this.GlobalConf.WordFile)
	res:=[]*model.Word{}

	for _,v:=range Chapter {
		index := sort.SearchStrings(request.Section,v.Title)

		if index < len(request.Section) && request.Section[index]==v.Title {
			res=append(res,v.Section...)
		}
	}

	return &response.GetWordResponse{Word:res}
}

func (this *WordController) POST_AddWord(request *request.AddWordRequest) *bootstrapmodel.BaseResponse{
	wordtool.AddWord(&model.Chapter{Title:request.Title,Section: request.Word},this.GlobalConf.WordFile)
	return &bootstrapmodel.BaseResponse{Result:true}
}