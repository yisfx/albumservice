package utils

import (
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
			log.Error("err %s", err)
		}
	default:
		{ // 非运行时错误
			log.Error("err %s", err)
		}
	}
}
