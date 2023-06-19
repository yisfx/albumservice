package main

import (
	"albumservice/albumtool"
	"albumservice/controller"
	"albumservice/framework/bootstrap"
	"albumservice/framework/configTool"
	"albumservice/framework/redisTool"
	"albumservice/framework/utils"
	"albumservice/interceptor"
	"fmt"
	"net/http"
	"strconv"
	"runtime"
	"strings"

	log "github.com/skoo87/log4go"
)

func main() {
	defer utils.HanderError()
	if err := log.SetupLogWithConf("./conf/logger.json"); err != nil {
		panic(err)
	}
	defer log.Close()

	conf := configTool.ReadSysConf()
	globalConf := configTool.ReadGlobalConf((conf.GlobalConfig))
	fmt.Printf("global conf:%#v \nglobalConf:%#v\n", conf, globalConf)

	// s := albumUtils.EncryptImageUri("aaa", "p", "max")
	// fmt.Println(s)
	// fmt.Printf("%#v \n", albumUtils.DecryptImageUri(s))
	// fmt.Println("***************************")
	// utils.DesDemo("abcdefd")
	// fmt.Println("***************************")
	// return


	if !strings.EqualFold(runtime.GOOS ,"windows") {
		redisClient := redisTool.RedisConnect(globalConf.Redis.Port, globalConf.Redis.Pwd)
		if redisClient.PoolStats().TotalConns < 1 {
			log.Error("redis connect failure")
			return
		}
		defer redisClient.Close()
	}

	bootstrap.SetConfig(*conf, *globalConf)

	bootstrap.AddInterceptor(interceptor.NewDemoInterceptor())

	bootstrap.Bootstrap(
		*bootstrap.NewControllerData("Manage", controller.NewAlbumManageController()),
		*bootstrap.NewControllerData("Login", controller.NewLoginController()),
		*bootstrap.NewControllerData("Demo", controller.NewDemoController()),
		*bootstrap.NewControllerData("Entry", controller.NewEntryController()),
		*bootstrap.NewControllerData("Word", controller.NewWordController()),
	)

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("hello world1"))
	})
	port := "0.0.0.0:" + strconv.Itoa(conf.Port)
	fmt.Println("listen at " + port)
	go albumtool.Out()
	http.ListenAndServe(port, nil)
}
