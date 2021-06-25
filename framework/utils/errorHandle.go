package utils

import (
	"fmt"
	"runtime"

	log "github.com/skoo87/log4go"
)

func ErrorHandler() {
	err := recover()
	if err == nil {
		return
	}
	switch err.(type) {
	case runtime.Error:
		{ // 运行时错误
			fmt.Println("err %s", err)
			log.Error("err %s", err)
		}
	default:
		{ // 非运行时错误
			fmt.Println("err %s", err)
			log.Error("err %s", err)
		}
	}
}
