package controller

import (
	"path"

	"albumservice/albumtool"
	"albumservice/albumtool/albumUtils"
	"albumservice/albumtool/constfield"
	"albumservice/albumtool/loginHelper"
	"albumservice/filter"
	"albumservice/framework/bootstrap"
	"albumservice/framework/bootstrapmodel"
	"albumservice/framework/fxfilter"
	m "albumservice/model"
	"albumservice/model/request"
	"albumservice/model/response"
)

type AlbumController struct {
	SysConfig   bootstrapmodel.SysConf
	GlobalConf  bootstrapmodel.GlobalConf
	AlbumHelper albumtool.AlbumHelper
	Context     *bootstrapmodel.Context
	IsLogin     *bool
}

func NewAlbumManageController() bootstrap.BaseController {
	o := &AlbumController{}
	return o
}
func (controller AlbumController) GetFilterMapping() fxfilter.FilterMapping {
	mapping := fxfilter.FilterMapping{}

	mapping["AddAlbum"] = fxfilter.FilterFuncList{filter.LoginFilter}
	mapping["BuildAlbumImage"] = fxfilter.FilterFuncList{filter.LoginFilter}
	mapping["BuildAllAlbum"] = fxfilter.FilterFuncList{filter.LoginFilter}
	mapping["BuildAllYears"] = fxfilter.FilterFuncList{filter.LoginFilter}
	mapping["BuildPicForAlbum"] = fxfilter.FilterFuncList{filter.LoginFilter}
	mapping["DeleteAlbum"] = fxfilter.FilterFuncList{filter.LoginFilter}
	mapping["DeleteAlbumPic"] = fxfilter.FilterFuncList{filter.LoginFilter}
	mapping["UploadImage"] = fxfilter.FilterFuncList{filter.LoginFilter}
	mapping["UploadImagePart"] = fxfilter.FilterFuncList{filter.LoginFilter}

	mapping["GetAlbumList"] = fxfilter.FilterFuncList{}
	mapping["GetAlbumListByYear"] = fxfilter.FilterFuncList{}
	mapping["GetAlbumPicList"] = fxfilter.FilterFuncList{}
	mapping["GetAllYears"] = fxfilter.FilterFuncList{}

	return mapping
}

// func (controller *AlbumController) Post_GetAlbumList() response.AlbumListResponse {
// 	albumList := controller.AlbumHelper.GetAlbumList()
// 	result := response.AlbumListResponse{}
// 	result.BaseResponse.Result = true
// 	result.AlbumList = albumList
// 	return result
// }

func (controller *AlbumController) Post_AddAlbum(r *request.AddAlbumRequest) *response.AddAlbumResponse {
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

// func (controller *AlbumController) Post_UploadImage(r *request.UploadPictureRequest) *bootstrapmodel.BaseResponse {
// 	result := new(bootstrapmodel.BaseResponse)
// 	album := controller.AlbumHelper.GetAlbum(r.AlbumName)
// 	for _, pic := range album.PicList {
// 		if strings.EqualFold(pic.Name, r.PictureName) {
// 			result.Result = false
// 			result.ErrorMessage = "picture exist"
// 			return result
// 		}
// 	}
// 	controller.AlbumHelper.AddAlbumPicture(album, r.PictureName)
// 	result.Result = true
// 	return result
// }

func (controller *AlbumController) Post_BuildAlbumImage(r *request.GetAlbumPicListRequest) *bootstrapmodel.BaseResponse {
	result := new(bootstrapmodel.BaseResponse)
	if controller.AlbumHelper.ExistsAlbum(r.AlbumName) {
		go albumtool.In(r.AlbumName)
		result.Result = true
	} else {
		result.Result = false
		result.ErrorMessage = "hasn't this Album"
	}
	return result
}

func (controller *AlbumController) Post_DeleteAlbum(r *request.DeleteAlbumRequest) *bootstrapmodel.BaseResponse {
	result := new(bootstrapmodel.BaseResponse)
	album := controller.AlbumHelper.GetAlbum(r.AlbumName)
	for _, pic := range album.PicList {
		controller.AlbumHelper.DeleteAlbumPic(album, pic.Name, m.DeleteImage)
	}
	controller.AlbumHelper.DeleteAlbum(album)
	result.Result = true
	return result
}

func (controller *AlbumController) Post_DeleteAlbumPic(r *request.DeleteAlbumPicRequest) *bootstrapmodel.BaseResponse {
	result := new(bootstrapmodel.BaseResponse)
	album := controller.AlbumHelper.GetAlbum(r.AlbumName)
	controller.AlbumHelper.DeleteAlbumPic(album, r.PicName, r.DeleteType)
	result.Result = true
	return result
}

func (controller *AlbumController) Post_UploadImagePart(r *request.PicturePartUploadRequest) *bootstrapmodel.BaseResponse {
	controller.AlbumHelper.CacheUploadImage(r.AlbumName, r.PictureName, r.PartIndex, r.Value)
	if r.IsLastPart {
		controller.AlbumHelper.BuildCacheUploadImage(r.AlbumName, r.PictureName, r.PartIndex)
	}
	result := new(bootstrapmodel.BaseResponse)
	result.Result = true
	return result
}

func (controller *AlbumController) Post_BuildAllAlbum() *bootstrapmodel.BaseResponse {
	controller.AlbumHelper.BuildAlbumList(controller.GlobalConf.AlbumPath)
	result := bootstrapmodel.BaseResponse{}
	result.Result = true
	return &result
}

func (controller *AlbumController) Post_GetAllYears() *response.GetAllYearsResponse {
	result := &response.GetAllYearsResponse{}
	result.AllYears = controller.AlbumHelper.GetAllYears()
	result.Result = true
	return result
}

func (controller *AlbumController) Post_BuildAllYears() *bootstrapmodel.BaseResponse {
	controller.AlbumHelper.BuildAllYears()
	result := &bootstrapmodel.BaseResponse{}
	result.Result = true
	return result
}

func (controller *AlbumController) Post_BuildPicForAlbum() *bootstrapmodel.BaseResponse {
	albumList := controller.AlbumHelper.GetAlbumList()
	for _, album := range albumList {
		controller.AlbumHelper.BuildPicForAlbum(album)
	}
	result := new(bootstrapmodel.BaseResponse)
	result.Result = true
	return result
}

func (controller *AlbumController) Post_GetAlbumListByYear(r *request.GetYearAlbumListRequest) *response.AlbumListResponse {
	result := &response.AlbumListResponse{}
	result.AlbumList = controller.AlbumHelper.GetAlbumListByYear(r.Year)

	///if not login encrypt image
	if !loginHelper.ValidateLoginStatus(controller.Context.GetParam(constfield.Header_Login_Token_Key)) {
		for _, album := range result.AlbumList {
			album.Cover = albumUtils.EncryptImageUri(album.Name, album.Cover, "max")
			album.CNName = albumUtils.EncryptAlbumName(album.Name)
			album.Name = album.CNName
			album.PicList = nil
		}
	}
	result.Result = true
	return result
}
