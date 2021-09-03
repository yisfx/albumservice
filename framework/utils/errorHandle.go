package utils

import (
	"fmt"
	"runtime"
	"runtime/debug"

	log "github.com/skoo87/log4go"
)

func buildMsg(prefix string, err interface{}) string {
	return fmt.Sprintf("\n [--------------------------------------------------------------------------------]\n %v:%#v \n  %s \n [--------------------------------------------------------------------------------] \n", prefix, err, debug.Stack())
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
			fmt.Println(msg)
			log.Error(msg)
		}
	default:
		{ // 非运行时错误
			msg := buildMsg("error", err)
			fmt.Println(msg)
			log.Error(msg)
		}
	}
	return err
}
