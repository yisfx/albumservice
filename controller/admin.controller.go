package controller

import (
	"encoding/json"
	"net/http"

	"../helper"
	"../model"
)

type AlbumManage struct {
	SysConf model.SysConf
}

func NewAlbumManageController(SysConf model.SysConf) *AlbumManage {
	o := &AlbumManage{}
	o.SysConf = SysConf
	return o
}

func (controller *AlbumManage) GeAlbumList(response http.ResponseWriter, request *http.Request) {
	albumHelper := helper.AlbumHelper{}
	albumList := albumHelper.BuildAlbumList(controller.SysConf.AlbumPath)
	resp, err := json.Marshal(albumList)
	if err == nil {
		response.Write(resp)
	} else {
		response.Write([]byte("500"))
	}
}

func (controller *AlbumManage) AddAlbum(response http.ResponseWriter, request *http.Request) {
	a := &model.Album{}
	json.Unmarshal(helper.ReadBody(request.Body), a)
	albumHelper := helper.AlbumHelper{}
	if albumHelper.ExistsAlbum(a.Name, controller.SysConf.AlbumPath) {
		response.Write([]byte("exists"))
	} else {
		albumHelper.CreateAlbum(*a)
		response.Write([]byte("success"))
	}

}
