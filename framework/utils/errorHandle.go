package utils

import (
	"fmt"
	"runtime"
	"runtime/debug"

	log "github.com/skoo87/log4go"
)

func buildMsg(prefix string, err interface{}) string {
	_, filename, line, _ := runtime.Caller(2)
	return fmt.Sprintf("\n--------------------------------------------------------------------------------\n [%v]:\n%v line:%v\n%#v \n %s \n[--------------------------------------------------------------------------------] \n",
		prefix, filename, line, err, debug.Stack())
}

func HanderError() interface{} {
	err := recover()
	if err == nil {
		return nil
	}

	switch err.(type) {
	case runtime.Error:
		{ // 运行时错误
			msg := buildMsg("runtime error", err)
			// fmt.Println(msg)
			log.Error(msg)
		}
	default:
		{ // 非运行时错误
			msg := buildMsg("error", err)
			// fmt.Println(msg)
			log.Error(msg)
		}
	}
	return err
}

func ProcessError(err interface{}) {
	if err == nil {
		return
	}

	switch err.(type) {
	case runtime.Error:
		{ // 运行时错误
			msg := buildMsg("runtime error", err)
			// fmt.Println(msg)
			log.Error(msg)
		}
	default:
		{ // 非运行时错误
			msg := buildMsg("error", err)
			// fmt.Println(msg)
			log.Error(msg)
		}
	}
}
