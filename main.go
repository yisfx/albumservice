package main

import (
	"albumservice/albumtool"
	"albumservice/controller"
	"albumservice/framework/bootstrap"
	"albumservice/framework/configTool"
	"albumservice/framework/redisTool"
	"albumservice/framework/model"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	conf := *configTool.ReadSysConf()
	globalConf := *configTool.ReadGlobalConf((conf.GlobalConfig))
	fmt.Println(conf, globalConf)

	redisTool.RedisConnect(globalConf.Redis.Port, globalConf.Redis.Pwd)
	// framework.ExampleClient_Hash()
	// framework.ExampleClient_Set()
	// framework.ExampleClient_SortSet()
	// framework.ExampleClient_HyperLogLog()
	// framework.ExampleClient_CMD()
	// framework.ExampleClient_Scan()
	// framework.ExampleClient_Tx()
	// framework.ExampleClient_Script()
	// framework.ExampleClient_PubSub()
	manageController := controller.NewAlbumManageController(conf, globalConf)

	bootstrap.Bootstrap(
		*model.NewControllerData("Manage", &manageController),
	)

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("hello world"))
	})
	port := "0.0.0.0:" + strconv.Itoa(conf.Port)
	fmt.Println("listen at " + port)
	go albumtool.Out()
	http.ListenAndServe(port, nil)
}
