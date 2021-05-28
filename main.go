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

	framework.RedisConnect(globalConf.Redis.Port, globalConf.Redis.Pwd)

	framework.ExampleClient_String()
	framework.ExampleClient_List()
	framework.ExampleClient_Hash()
	framework.ExampleClient_Set()
	framework.ExampleClient_SortSet()
	framework.ExampleClient_HyperLogLog()
	framework.ExampleClient_CMD()
	framework.ExampleClient_Scan()
	framework.ExampleClient_Tx()
	framework.ExampleClient_Script()
	framework.ExampleClient_PubSub()

	manageController := controller.NewAlbumManageController(conf, globalConf)
	framework.Bootstrap(&manageController, "/Manage/", manageController.Process)

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("hello world"))
	})
	port := "0.0.0.0:" + strconv.Itoa(conf.Port)
	fmt.Println("listen at " + port)
	go albumtool.Out()
	http.ListenAndServe(port, nil)
}
