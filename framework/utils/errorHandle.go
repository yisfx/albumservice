package utils

import (
	"fmt"
	"runtime"

	log "github.com/skoo87/log4go"
)

func HanderError(funcName string) {
	err := recover()
	if err == nil {
		return
	}
	switch err.(type) {
	case runtime.Error:
		{ // 运行时错误
			fmt.Println(fmt.Sprintf("%s: %s", funcName, err))
			log.Error(fmt.Sprintf("%s: %s", funcName, err))
		}
	default:
		{ // 非运行时错误
			fmt.Println(fmt.Sprintf("%s: %s", funcName, err))
			log.Error(fmt.Sprintf("%s: %s", funcName, err))
		}
	}
}

func ErrorHandler() {
	HanderError("old")
}
