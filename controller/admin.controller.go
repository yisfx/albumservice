package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"reflect"

	"albumservice/albumtool"
	"albumservice/framework"
	model "albumservice/model"
	requestModel "albumservice/model/request"
	responseModel "albumservice/model/response"
)

type AlbumManage struct {
	SysConf   model.SysConf
	RouterMap map[string]reflect.Type
}

func NewAlbumManageController(SysConf model.SysConf) *AlbumManage {
	o := &AlbumManage{}
	o.SysConf = SysConf
	o.RouterMap = make(map[string]reflect.Type)
	controller := reflect.TypeOf(o)
	for i := 0; i < controller.NumMethod(); i++ {
		m := controller.Method(i)
		if m.Type.NumIn() > 1 {
			o.RouterMap[m.Name] = m.Type.In(1)
		} else {
			o.RouterMap[m.Name] = reflect.TypeOf(nil)
		}
	}
	return o
}

func (controller *AlbumManage) GetAlbumList(res http.ResponseWriter, request *http.Request) {
	albumHelper := albumtool.AlbumHelper{}
	albumList := albumHelper.BuildAlbumList(controller.SysConf.AlbumPath)
	result := new(responseModel.AlbumListResponse)
	result.BaseResponse.Result = true
	result.AlbumList = albumList
	resp, err := json.Marshal(result)
	if err == nil {
		res.Write(resp)
	} else {
		res.Write([]byte("500"))
	}
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
