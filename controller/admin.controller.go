package controller

import (
	"path"
	"strings"

	"albumservice/albumtool"
	"albumservice/framework/model"
	"albumservice/framework/utils"
	m "albumservice/model"
	"albumservice/model/request"
	"albumservice/model/response"
)

type AlbumManage struct {
	SysConfig  model.SysConf
	GlobalConf model.GlobalConf
	model.BaseController
	albumHelper *albumtool.AlbumHelper
}

func NewAlbumManageController() model.BaseController {
	o := &AlbumManage{}
	return o
}

func (controller *AlbumManage) Post_GetAlbumList() response.AlbumListResponse {
	defer utils.ErrorHandler()
	albumList := controller.albumHelper.GetAlbumList()
	result := response.AlbumListResponse{}
	result.BaseResponse.Result = true
	result.AlbumList = albumList
	return result
}

func (controller *AlbumManage) Post_AddAlbum(r *request.AddAlbumRequest) *response.AddAlbumResponse {
	defer utils.ErrorHandler()
	a := r.Album
	result := new(response.AddAlbumResponse)
	a.Path = path.Join(controller.GlobalConf.AlbumPath, a.Name)
	if controller.albumHelper.ExistsAlbum(a.Name) {
		///edit
		controller.albumHelper.EditAlbum(a)
		result.BaseResponse.Result = true
	} else {
		///add
		controller.albumHelper.CreateAlbum(a)
		result.BaseResponse.Result = true
	}
	return result
}

func (controller *AlbumManage) Post_GetAlbumPicList(r *request.GetAlbumPicListRequest) *response.GetAlbumPicListResponse {
	defer utils.ErrorHandler()
	result := new(response.GetAlbumPicListResponse)
	if !controller.albumHelper.ExistsAlbum(r.AlbumName) {
		result.BaseResponse.Result = false
		result.BaseResponse.ErrorMessage = "hasn't this Album"
	} else {
		result.BaseResponse.Result = true
		result.Album = *controller.albumHelper.GetAlbum(r.AlbumName)
	}

	return result
}

func (controller *AlbumManage) Post_UploadImage(r *request.UploadPictureRequest) *response.BaseResponse {
	defer utils.ErrorHandler()
	result := new(response.BaseResponse)
	album := controller.albumHelper.GetAlbum(r.AlbumName)
	for _, pic := range album.PicList {
		if strings.EqualFold(pic.Name, r.PictureName) {
			result.Result = false
			result.ErrorMessage = "picture exist"
			return result
		}
	}
	controller.albumHelper.AddAlbumPicture(album, r.PictureName)
	result.Result = true
	return result
}

func (controller *AlbumManage) Post_BuildAlbumImage(r *request.GetAlbumPicListRequest) *response.BaseResponse {
	defer utils.ErrorHandler()
	result := new(response.BaseResponse)
	if controller.albumHelper.ExistsAlbum(r.AlbumName) {
		go albumtool.In(r.AlbumName)
		result.Result = true
	} else {
		result.Result = false
		result.ErrorMessage = "hasn't this Album"
	}
	return result
}

func (controller *AlbumManage) Post_DeleteAlbum(r *request.DeleteAlbumRequest) *response.BaseResponse {
	defer utils.ErrorHandler()
	result := new(response.BaseResponse)
	album := controller.albumHelper.GetAlbum(r.AlbumName)
	for _, pic := range album.PicList {
		controller.albumHelper.DeleteAlbumPic(album, pic.Name, m.DeleteImage)
	}
	controller.albumHelper.DeleteAlbum(album)
	result.Result = true
	return result
}

func (controller *AlbumManage) Post_DeleteAlbumPic(r *request.DeleteAlbumPicRequest) *response.BaseResponse {
	defer utils.ErrorHandler()
	result := new(response.BaseResponse)
	album := controller.albumHelper.GetAlbum(r.AlbumName)
	controller.albumHelper.DeleteAlbumPic(album, r.PicName, r.DeleteType)
	result.Result = true
	return result
}

func (controller *AlbumManage) Post_UploadImagePart(r *request.PicturePartUploadRequest) *response.BaseResponse {
	defer utils.ErrorHandler()
	controller.albumHelper.CacheUploadImage(r.AlbumName, r.PictureName, r.PartIndex, r.Value)
	if r.IsLastPart {
		controller.albumHelper.BuildCacheUploadImage(r.AlbumName, r.PictureName, r.PartIndex)
	}
	result := new(response.BaseResponse)
	result.Result = true
	return result
}

func (controller *AlbumManage) Post_BuildAllAlbum() *response.BaseResponse {
	defer utils.ErrorHandler()

	controller.albumHelper.BuildAlbumList(controller.GlobalConf.AlbumPath)
	result := response.BaseResponse{}
	result.Result = true
	return &result
}

func (controller *AlbumManage) Post_GetAllYears() *response.GetAllYearsResponse {
	defer utils.ErrorHandler()
	result := &response.GetAllYearsResponse{}
	result.AllYears = controller.albumHelper.GetAllYears()
	result.Result = true
	return result
}

func (controller *AlbumManage) Post_BuildAllYears() *response.BaseResponse {
	defer utils.ErrorHandler()
	controller.albumHelper.BuildAllYears()
	result := &response.BaseResponse{}
	result.Result = true
	return result
}

func (controller *AlbumManage) Post_BuildPicForAlbum() *response.BaseResponse {
	defer utils.ErrorHandler()
	albumList := controller.albumHelper.GetAlbumList()
	for _, album := range albumList {
		controller.albumHelper.BuildPicForAlbum(album)
	}
	result := new(response.BaseResponse)
	result.Result = true
	return result
}

func (controller *AlbumManage) Post_GetAlbumListByYear(r *request.GetYearAlbumListRequest) *response.AlbumListResponse {
	result := &response.AlbumListResponse{}
	result.AlbumList = controller.albumHelper.GetAlbumListByYear(r.Year)
	result.Result = true
	return result
}
