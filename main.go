package main

import (
	"fmt"
	"net/http"
	"strconv"

	"albumservice/albumtool"
	"albumservice/controller"
	"albumservice/framework"
)

func main() {
	conf := *framework.ReadSysConf()
	globalConf := *framework.ReadGlobalConf((conf.GlobalConfig))
	fmt.Println(conf, globalConf)
	manageController := controller.NewAlbumManageController(conf, globalConf)
	framework.Bootstrap(&manageController, "/Manage/", manageController.Process)

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("hello world"))
	})
	port := ":" + strconv.Itoa(conf.Port)
	fmt.Println("listen at " + port)
	go albumtool.Out()
	http.ListenAndServe(port, nil)
}
