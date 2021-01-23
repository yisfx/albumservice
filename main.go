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
	fmt.Println(conf)
	manageController := controller.NewAlbumManageController(conf)

	http.HandleFunc("/Manage/AlbumList", manageController.GeAlbumList)
	http.HandleFunc("/Manage/AddAlbum", manageController.AddAlbum)
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("hello world"))
	})
	port := ":" + strconv.Itoa(conf.Port)
	fmt.Println("listen at " + port)
	http.ListenAndServe(port, nil)
}
