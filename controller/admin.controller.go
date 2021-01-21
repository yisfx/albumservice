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
