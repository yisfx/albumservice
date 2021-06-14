package controller

import (
	"path"
	"strings"

	"albumservice/albumtool"
	model "albumservice/model"
	"albumservice/model/request"
	"albumservice/model/response"
)

type AlbumManage struct {
	SysConfig  model.SysConf
	RouterList map[string]model.RouterMap
	GlobalConf model.GlobalConf
	model.BaseController
}

func NewAlbumManageController(SysConf model.SysConf, GlobalConf model.GlobalConf) model.BaseController {
	o := &AlbumManage{}
	o.SysConfig = SysConf
	o.GlobalConf = GlobalConf
	return o
}

func (controller *AlbumManage) Post_GetAlbumList() response.AlbumListResponse {
	albumHelper := albumtool.AlbumHelper{}
	albumList := albumHelper.GetAlbumList()
	result := response.AlbumListResponse{}
	result.BaseResponse.Result = true
	result.AlbumList = albumList
	return result
}

func (controller *AlbumManage) Post_AddAlbum(r *request.AddAlbumRequest) *response.AddAlbumResponse {
	a := r.Album
	albumHelper := albumtool.NewAlbumHelper()
	result := new(response.AddAlbumResponse)
	a.Path = path.Join(controller.GlobalConf.AlbumPath, a.Name)
	if albumHelper.ExistsAlbum(a.Name) {
		///edit
		albumHelper.EditAlbum(a)
		result.BaseResponse.Result = true
	} else {
		///add
		albumHelper.CreateAlbum(a)
		result.BaseResponse.Result = true
	}
	return result
}

func (controller *AlbumManage) Post_GetAlbumPicList(r *request.GetAlbumPicListRequest) *response.GetAlbumPicListResponse {
	albumHelper := albumtool.NewAlbumHelper()
	result := new(response.GetAlbumPicListResponse)
	if !albumHelper.ExistsAlbum(r.AlbumName) {
		result.BaseResponse.Result = false
		result.BaseResponse.ErrorMessage = "hasn't this Album"
	} else {
		result.BaseResponse.Result = true
		result.Album = *albumHelper.GetAlbum(r.AlbumName)
	}

	return result
}

func (controller *AlbumManage) Post_UploadImage(r *request.UploadPictureRequest) *response.BaseResponse {
	albumHelper := albumtool.NewAlbumHelper()
	result := new(response.BaseResponse)
	album := albumHelper.GetAlbum(r.AlbumName)
	for _, pic := range album.PicList {
		if strings.EqualFold(pic.Name, r.PictureName) {
			result.Result = false
			result.ErrorMessage = "picture exist"
			return result
		}
	}
	albumHelper.AddAlbumPicture(album, r.PictureName)
	result.Result = true
	return result
}

func (controller *AlbumManage) Post_BuildAlbumImage(r *request.GetAlbumPicListRequest) *response.BaseResponse {
	albumHelper := albumtool.NewAlbumHelper()
	result := new(response.BaseResponse)
	if albumHelper.ExistsAlbum(r.AlbumName) {
		go albumtool.In(r.AlbumName)
		result.Result = true
	} else {
		result.Result = false
		result.ErrorMessage = "hasn't this Album"
	}
	return result
}

func (controller *AlbumManage) Post_DeleteAlbumPic(r *request.DeleteAlbumPicRequest) *response.BaseResponse {
	albumHelper := albumtool.NewAlbumHelper()
	result := new(response.BaseResponse)
	album := albumHelper.GetAlbum(r.AlbumName)
	albumHelper.DeleteAlbumPic(album, r.PicName, r.DeleteType)
	result.Result = true
	return result
}

func (controller *AlbumManage) Post_UploadImagePart(r *request.PicturePartUploadRequest) *response.BaseResponse {
	albumHelper := &albumtool.AlbumHelper{}
	albumHelper.CacheUploadImage(r.AlbumName, r.PictureName, r.PartIndex, r.Value)
	if r.IsLastPart {
		albumHelper.BuildCacheUploadImage(r.AlbumName, r.PictureName, r.PartIndex)
	}
	result := new(response.BaseResponse)
	result.Result = true
	return result
}

func (controller *AlbumManage) Post_BuildAllAlbum() *response.BaseResponse {
	albumHelper := &albumtool.AlbumHelper{}
	albumHelper.BuildAlbumList(controller.GlobalConf.AlbumPath)
	result := response.BaseResponse{}
	result.Result = true
	return &result
}
