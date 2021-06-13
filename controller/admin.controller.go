package controller

import (
	"path"
	"strings"

	"albumservice/albumtool"
	model "albumservice/model"
	requestModel "albumservice/model/request"
	responseModel "albumservice/model/response"
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

func (controller *AlbumManage) Post_GetAlbumList() responseModel.AlbumListResponse {
	albumHelper := albumtool.AlbumHelper{}
	albumList := albumHelper.GetAlbumList()
	result := responseModel.AlbumListResponse{}
	result.BaseResponse.Result = true
	result.AlbumList = albumList
	return result
}

func (controller *AlbumManage) Post_AddAlbum(r *requestModel.AddAlbumRequest) *responseModel.AddAlbumResponse {
	a := r.Album
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.AddAlbumResponse)
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

func (controller *AlbumManage) Post_GetAlbumPicList(r *requestModel.GetAlbumPicListRequest) *responseModel.GetAlbumPicListResponse {
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.GetAlbumPicListResponse)
	if !albumHelper.ExistsAlbum(r.AlbumName) {
		result.BaseResponse.Result = false
		result.BaseResponse.ErrorMessage = "hasn't this Album"
	} else {
		result.BaseResponse.Result = true
		result.Album = *albumHelper.GetAlbum(r.AlbumName)
	}

	return result
}

func (controller *AlbumManage) Post_UploadImage(r *requestModel.UploadPictureRequest) *responseModel.BaseResponse {
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.BaseResponse)
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

func (controller *AlbumManage) Post_BuildAlbumImage(r *requestModel.GetAlbumPicListRequest) *responseModel.BaseResponse {
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.BaseResponse)
	if albumHelper.ExistsAlbum(r.AlbumName) {
		go albumtool.In(r.AlbumName)
		result.Result = true
	} else {
		result.Result = false
		result.ErrorMessage = "hasn't this Album"
	}
	return result
}

func (controller *AlbumManage) Post_DeleteAlbumPic(r *requestModel.DeleteAlbumPicRequest) *responseModel.BaseResponse {
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.BaseResponse)
	albumHelper.DeleteAlbumPic(path.Join(controller.GlobalConf.AlbumPath, r.AlbumName), r.PicName, r.DeleteType)
	result.Result = true
	return result
}

func (controller *AlbumManage) Post_UploadImagePart(r *requestModel.PicturePartUploadRequest) {
	albumHelper := &albumtool.AlbumHelper{}
	albumHelper.CacheUploadImage(r.AlbumName, r.PictureName, r.PartIndex, r.Value)
	if r.IsLastPart {
		albumHelper.BuildCacheUploadImage(r.AlbumName, r.PictureName, r.PartIndex)
	}

}
