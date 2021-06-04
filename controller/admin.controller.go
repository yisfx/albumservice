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
	GlobalConf model.GlobalConf
	model.BaseController
}

func NewAlbumManageController(SysConf model.SysConf, GlobalConf model.GlobalConf) model.BaseController {
	o := &AlbumManage{}
	o.SysConfig = SysConf
	o.GlobalConf = GlobalConf
	return o
}

func (controller *AlbumManage) Process(resp http.ResponseWriter, request *http.Request) {
	urls := strings.Split(request.URL.Path, "/")

	route, isok := controller.RouterList[urls[2]]
	if !isok {
		var result404, _ = json.Marshal(404)
		resp.Write(result404)
		return
	}
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

func (controller *AlbumManage) Post_BuildAlbumImage(r *requestModel.GetAlbumPicListRequest) *responseModel.BaseResponse {
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.BaseResponse)
	if albumHelper.ExistsAlbum(r.AlbumName) {
		go albumtool.In(path.Join(controller.GlobalConf.AlbumPath, r.AlbumName))
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
