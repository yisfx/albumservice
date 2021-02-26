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
	SysConf   model.SysConf
	RouterMap map[string]model.RouterMap
}

func NewAlbumManageController(SysConf model.SysConf) *AlbumManage {
	o := &AlbumManage{}
	o.SysConf = SysConf
	o.RouterMap = make(map[string]model.RouterMap)
	controller := reflect.TypeOf(o)
	for i := 0; i < controller.NumMethod(); i++ {

		m := controller.Method(i)
		if strings.EqualFold(m.Name, "process") {
			continue
		}
		route := model.RouterMap{}
		route.Controller = reflect.ValueOf(m)
		if m.Type.NumIn() > 1 {
			route.ArgType = m.Type.In(1)
		} else {
			route.ArgType = nil
		}
		o.RouterMap[m.Name] = route

	}
	return o
}

func (controller *AlbumManage) Process(res http.ResponseWriter, request *http.Request) {
	urls := strings.Split(request.URL.Path, "/")

	route := controller.RouterMap[urls[2]]
	if route.ArgType == nil {
		fmt.Println(route.Controller.Call(nil))
	}
	res.Write([]byte("123123"))
}

func (controller *AlbumManage) GetAlbumList() responseModel.AlbumListResponse {
	albumHelper := albumtool.AlbumHelper{}
	albumList := albumHelper.BuildAlbumList(controller.SysConf.AlbumPath)
	result := responseModel.AlbumListResponse{}
	result.BaseResponse.Result = true
	result.AlbumList = albumList
	return result
}

func (controller *AlbumManage) AddAlbum(resp http.ResponseWriter, request *http.Request) {

	a := &model.Album{}
	json.Unmarshal(framework.ReadBody(request.Body), a)
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.AddAlbumResponse)
	a.Path = path.Join(controller.SysConf.AlbumPath, a.Name)
	if albumHelper.ExistsAlbum(a.Name, controller.SysConf.AlbumPath) {
		///edit
		albumHelper.EditAlbum(*a)
		result.BaseResponse.Result = true

		// result.BaseResponse.ErrorMessage = "album exists"
	} else {
		///add
		albumHelper.CreateAlbum(*a)
		result.BaseResponse.Result = true
	}
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
	resp.Write(r)
}

func (controller *AlbumManage) GetAlbumPicList(resp http.ResponseWriter, request *http.Request) {
	r := &requestModel.GetAlbumPicListRequest{}
	json.Unmarshal(framework.ReadBody(request.Body), r)
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.GetAlbumPicListResponse)
	if !albumHelper.ExistsAlbum(r.AlbumName, controller.SysConf.AlbumPath) {
		result.BaseResponse.Result = false
		result.BaseResponse.ErrorMessage = "hasn't this Album"
	} else {
		result.BaseResponse.Result = true
		result.Album = albumHelper.GetAlbum(path.Join(controller.SysConf.AlbumPath, r.AlbumName))
	}

	b, _ := json.Marshal(result)
	resp.Write(b)
}

func (controller *AlbumManage) BuildAlbumImage(resp http.ResponseWriter, request *http.Request) {
	r := &requestModel.GetAlbumPicListRequest{}
	json.Unmarshal(framework.ReadBody(request.Body), r)
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.BaseResponse)
	if albumHelper.ExistsAlbum(r.AlbumName, controller.SysConf.AlbumPath) {
		go albumtool.In(path.Join(controller.SysConf.AlbumPath, r.AlbumName))
		result.Result = true
	} else {
		result.Result = false
		result.ErrorMessage = "hasn't this Album"
	}
	b, _ := json.Marshal(result)
	resp.Write(b)
}

func (controller *AlbumManage) DeleteAlbumPic(resp http.ResponseWriter, request *http.Request) {
	r := &requestModel.DeleteAlbumPicRequest{}
	json.Unmarshal(framework.ReadBody(request.Body), r)
	albumHelper := albumtool.NewAlbumHelper()
	result := new(responseModel.BaseResponse)
	albumHelper.DeleteAlbumPic(path.Join(controller.SysConf.AlbumPath, r.AlbumName), r.PicName, r.DeleteType)
	result.Result = true
	b, _ := json.Marshal(result)
	resp.Write(b)
}
