package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"reflect"
	"strings"

	"albumservice/albumtool"
	"albumservice/framework"
	model "albumservice/model"
	requestModel "albumservice/model/request"
	responseModel "albumservice/model/response"
)

type AlbumManage struct {
	SysConfig  model.SysConf
	RouterList map[string]model.RouterMap
	model.BaseController
}

func NewAlbumManageController(SysConf model.SysConf) model.BaseController {
	o := &AlbumManage{}
	o.SysConfig = SysConf
	return o
}

func (controller *AlbumManage) Process(resp http.ResponseWriter, request *http.Request) {
	urls := strings.Split(request.URL.Path, "/")

	route := controller.RouterList[urls[2]]

	var result []reflect.Value
	if route.ArgType == nil {
		result = route.Controller.Call(nil)
	} else {
		a := reflect.New(route.ArgType).Interface()
		framework.MustJSONDecode(framework.ReadBody(request.Body), a)
		args := []reflect.Value{reflect.ValueOf(a)}
		result = route.Controller.Call(args)
	}

	r, err := json.Marshal(result[0].Interface())
	if err != nil {
		fmt.Println("err:", err)
	}
	resp.Write(r)
}

func (controller *AlbumManage) Post_GetAlbumList() responseModel.AlbumListResponse {
	albumHelper := albumtool.AlbumHelper{}
	albumList := albumHelper.BuildAlbumList(controller.SysConfig.AlbumPath)
	result := responseModel.AlbumListResponse{}
	result.BaseResponse.Result = true
	result.AlbumList = albumList
	return result
}

func (controller *AlbumManage) Post_AddAlbum(r *requestModel.AddAlbumRequest) *responseModel.AddAlbumResponse {
	a := r.Album
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.AddAlbumResponse)
	a.Path = path.Join(controller.SysConfig.AlbumPath, a.Name)
	if albumHelper.ExistsAlbum(a.Name, controller.SysConfig.AlbumPath) {
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
	if !albumHelper.ExistsAlbum(r.AlbumName, controller.SysConfig.AlbumPath) {
		result.BaseResponse.Result = false
		result.BaseResponse.ErrorMessage = "hasn't this Album"
	} else {
		result.BaseResponse.Result = true
		result.Album = albumHelper.GetAlbum(path.Join(controller.SysConfig.AlbumPath, r.AlbumName))
	}

	return result
}

func (controller *AlbumManage) Post_BuildAlbumImage(r *requestModel.GetAlbumPicListRequest) *responseModel.BaseResponse {
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.BaseResponse)
	if albumHelper.ExistsAlbum(r.AlbumName, controller.SysConfig.AlbumPath) {
		go albumtool.In(path.Join(controller.SysConfig.AlbumPath, r.AlbumName))
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
	albumHelper.DeleteAlbumPic(path.Join(controller.SysConfig.AlbumPath, r.AlbumName), r.PicName, r.DeleteType)
	result.Result = true
	return result
}
