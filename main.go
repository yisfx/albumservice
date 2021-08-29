package main

import (
	"albumservice/albumtool"
	"albumservice/controller"
	"albumservice/framework/bootstrap"
	"albumservice/framework/configTool"
	"albumservice/framework/model"
	"albumservice/framework/redisTool"
	"albumservice/framework/utils"
	"albumservice/interceptor"
	"fmt"
	"net/http"
	"strconv"

	log "github.com/skoo87/log4go"
)

func main() {
	defer utils.ErrorHandler()
	if err := log.SetupLogWithConf("./conf/logger.json"); err != nil {
		panic(err)
	}
	defer log.Close()

	conf := configTool.ReadSysConf()
	globalConf := configTool.ReadGlobalConf((conf.GlobalConfig))
	fmt.Println("global config:", conf, globalConf)
	// fmt.Println("***************************")
	// utils.DesDemo()
	// fmt.Println("***************************")
	// return
	redisClient := redisTool.RedisConnect(globalConf.Redis.Port, globalConf.Redis.Pwd)
	if redisClient.PoolStats().TotalConns < 1 {
		log.Error("redis connect failure")
		return
	}
	defer redisClient.Close()

	bootstrap.SetConfig(*conf, *globalConf)

	bootstrap.AddInterceptor(interceptor.NewLoginInterceptor())

	manageController := controller.NewAlbumManageController()

	bootstrap.Bootstrap(
		*model.NewControllerData("Manage", &manageController),
	)

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("hello world1"))
	})
	port := "0.0.0.0:" + strconv.Itoa(conf.Port)
	fmt.Println("listen at " + port)
	go albumtool.Out()
	http.ListenAndServe(port, nil)
}
