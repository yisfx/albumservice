package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"../helper"
	"../model"
	responseModel "../model/response"
)

type AlbumManage struct {
	SysConf model.SysConf
}

func NewAlbumManageController(SysConf model.SysConf) *AlbumManage {
	o := &AlbumManage{}
	o.SysConf = SysConf
	return o
}

func (controller *AlbumManage) GeAlbumList(res http.ResponseWriter, request *http.Request) {
	albumHelper := helper.AlbumHelper{}
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
	json.Unmarshal(helper.ReadBody(request.Body), a)
	albumHelper := helper.AlbumHelper{}
	result := new(responseModel.AddAlbumResponse)
	a.Path = path.Join(controller.SysConf.AlbumPath, a.Name)
	if albumHelper.ExistsAlbum(a.Name, controller.SysConf.AlbumPath) {
		result.BaseResponse.Result = false
		result.BaseResponse.ErrorMessage = "album exists"
	} else {
		albumHelper.CreateAlbum(*a)
		result.BaseResponse.Result = true
	}
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
	resp.Write(r)
}
