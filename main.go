package main

import (
	"fmt"
	"net/http"
	"strconv"

	"./controller"
	"./framework"
)

func main() {
	conf := *framework.ReadSysConf()
	manageController := &controller.AlbumManage{}
	http.HandleFunc("/Manage/AlbumList", manageController.GeAlbumList)

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("hello world"))
	})
	port := ":" + strconv.Itoa(conf.Port)
	fmt.Println("listen at " + port)
	http.ListenAndServe(port, nil)

}
