package controller

import (
	"net/http"
	"path"
	"strings"

	"albumservice/albumtool"
	"albumservice/filter"
	"albumservice/framework/bootstrap"
	"albumservice/framework/model"
	"albumservice/framework/utils"
	m "albumservice/model"
	"albumservice/model/request"
	"albumservice/model/response"
)

type AlbumController struct {
	SysConfig   model.SysConf
	GlobalConf  model.GlobalConf
	AlbumHelper albumtool.AlbumHelper
	Request     *http.Request
	IsLogin     *bool
}

func NewAlbumManageController() bootstrap.BaseController {
	o := &AlbumController{}
	return o
}
func (controller AlbumController) GetFilterMapping() bootstrap.FilterMapping {
	mapping := bootstrap.FilterMapping{}

	mapping["AddAlbum"] = bootstrap.FilterFuncList{filter.LoginFilter}
	mapping["BuildAlbumImage"] = bootstrap.FilterFuncList{filter.LoginFilter}
	mapping["BuildAllAlbum"] = bootstrap.FilterFuncList{filter.LoginFilter}
	mapping["BuildAllYears"] = bootstrap.FilterFuncList{filter.LoginFilter}
	mapping["BuildPicForAlbum"] = bootstrap.FilterFuncList{filter.LoginFilter}
	mapping["DeleteAlbum"] = bootstrap.FilterFuncList{filter.LoginFilter}
	mapping["DeleteAlbumPic"] = bootstrap.FilterFuncList{filter.LoginFilter}
	mapping["UploadImage"] = bootstrap.FilterFuncList{filter.LoginFilter}
	mapping["UploadImagePart"] = bootstrap.FilterFuncList{filter.LoginFilter}

	mapping["GetAlbumList"] = bootstrap.FilterFuncList{}
	mapping["GetAlbumListByYear"] = bootstrap.FilterFuncList{}
	mapping["GetAlbumPicList"] = bootstrap.FilterFuncList{}
	mapping["GetAllYears"] = bootstrap.FilterFuncList{}

	return mapping
}

func (controller *AlbumController) Post_GetAlbumList() response.AlbumListResponse {
	defer utils.ErrorHandler()
	albumList := controller.AlbumHelper.GetAlbumList()
	result := response.AlbumListResponse{}
	result.BaseResponse.Result = true
	result.AlbumList = albumList
	return result
}

func (controller *AlbumController) Post_AddAlbum(r *request.AddAlbumRequest) *response.AddAlbumResponse {
	defer utils.ErrorHandler()
	a := r.Album
	result := new(response.AddAlbumResponse)
	a.Path = path.Join(controller.GlobalConf.AlbumPath, a.Name)
	if controller.AlbumHelper.ExistsAlbum(a.Name) {
		///edit
		controller.AlbumHelper.EditAlbum(a)
		result.BaseResponse.Result = true
	} else {
		///add
		controller.AlbumHelper.CreateAlbum(a)
		result.BaseResponse.Result = true
	}
	return result
}

func (controller *AlbumController) Post_GetAlbumPicList(r *request.GetAlbumPicListRequest) *response.GetAlbumPicListResponse {
	defer utils.ErrorHandler()
	result := new(response.GetAlbumPicListResponse)
	if !controller.AlbumHelper.ExistsAlbum(r.AlbumName) {
		result.BaseResponse.Result = false
		result.BaseResponse.ErrorMessage = "hasn't this Album"
	} else {
		result.BaseResponse.Result = true
		result.Album = *controller.AlbumHelper.GetAlbum(r.AlbumName)
	}

	return result
}

func (controller *AlbumController) Post_UploadImage(r *request.UploadPictureRequest) *response.BaseResponse {
	defer utils.ErrorHandler()
	result := new(response.BaseResponse)
	album := controller.AlbumHelper.GetAlbum(r.AlbumName)
	for _, pic := range album.PicList {
		if strings.EqualFold(pic.Name, r.PictureName) {
			result.Result = false
			result.ErrorMessage = "picture exist"
			return result
		}
	}
	controller.AlbumHelper.AddAlbumPicture(album, r.PictureName)
	result.Result = true
	return result
}

func (controller *AlbumController) Post_BuildAlbumImage(r *request.GetAlbumPicListRequest) *response.BaseResponse {
	defer utils.ErrorHandler()
	result := new(response.BaseResponse)
	if controller.AlbumHelper.ExistsAlbum(r.AlbumName) {
		go albumtool.In(r.AlbumName)
		result.Result = true
	} else {
		result.Result = false
		result.ErrorMessage = "hasn't this Album"
	}
	return result
}

func (controller *AlbumController) Post_DeleteAlbum(r *request.DeleteAlbumRequest) *response.BaseResponse {
	defer utils.ErrorHandler()
	result := new(response.BaseResponse)
	album := controller.AlbumHelper.GetAlbum(r.AlbumName)
	for _, pic := range album.PicList {
		controller.AlbumHelper.DeleteAlbumPic(album, pic.Name, m.DeleteImage)
	}
	controller.AlbumHelper.DeleteAlbum(album)
	result.Result = true
	return result
}

func (controller *AlbumController) Post_DeleteAlbumPic(r *request.DeleteAlbumPicRequest) *response.BaseResponse {
	defer utils.ErrorHandler()
	result := new(response.BaseResponse)
	album := controller.AlbumHelper.GetAlbum(r.AlbumName)
	controller.AlbumHelper.DeleteAlbumPic(album, r.PicName, r.DeleteType)
	result.Result = true
	return result
}

func (controller *AlbumController) Post_UploadImagePart(r *request.PicturePartUploadRequest) *response.BaseResponse {
	defer utils.ErrorHandler()
	controller.AlbumHelper.CacheUploadImage(r.AlbumName, r.PictureName, r.PartIndex, r.Value)
	if r.IsLastPart {
		controller.AlbumHelper.BuildCacheUploadImage(r.AlbumName, r.PictureName, r.PartIndex)
	}
	result := new(response.BaseResponse)
	result.Result = true
	return result
}

func (controller *AlbumController) Post_BuildAllAlbum() *response.BaseResponse {
	defer utils.ErrorHandler()

	controller.AlbumHelper.BuildAlbumList(controller.GlobalConf.AlbumPath)
	result := response.BaseResponse{}
	result.Result = true
	return &result
}

func (controller *AlbumController) Post_GetAllYears() *response.GetAllYearsResponse {
	defer utils.ErrorHandler()
	result := &response.GetAllYearsResponse{}
	result.AllYears = controller.AlbumHelper.GetAllYears()
	result.Result = true
	return result
}

func (controller *AlbumController) Post_BuildAllYears() *response.BaseResponse {
	defer utils.ErrorHandler()
	controller.AlbumHelper.BuildAllYears()
	result := &response.BaseResponse{}
	result.Result = true
	return result
}

func (controller *AlbumController) Post_BuildPicForAlbum() *response.BaseResponse {
	defer utils.ErrorHandler()
	albumList := controller.AlbumHelper.GetAlbumList()
	for _, album := range albumList {
		controller.AlbumHelper.BuildPicForAlbum(album)
	}
	result := new(response.BaseResponse)
	result.Result = true
	return result
}

func (controller *AlbumController) Post_GetAlbumListByYear(r *request.GetYearAlbumListRequest) *response.AlbumListResponse {
	result := &response.AlbumListResponse{}
	result.AlbumList = controller.AlbumHelper.GetAlbumListByYear(r.Year)
	result.Result = true
	return result
}
