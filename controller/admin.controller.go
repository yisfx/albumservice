package controller

import (
	"net/http"
)

type AlbumManage struct{}

func (controller *AlbumManage) GeAlbumList(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(request.URL.Path))
}
