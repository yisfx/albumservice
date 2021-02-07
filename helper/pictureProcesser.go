package helper

import (
	"fmt"
	"time"

	"../framework"
	model "../model"
)

type PictureProcesser struct {
	sysConf model.SysConf
}

var message = make(chan string)

func (this *PictureProcesser) init() {
	this.sysConf = framework.ReadSysConf()
}

func (this *PictureProcesser) In(album string) {

	message <- album
	// time.Sleep(time.Second * 7)
	// for {
	// 	if time.Now().Second()%3 == 0 {
	// 		message <- "ininin"
	// 	}
	// 	time.Sleep(time.Second * 7)
	// }
}

func Out() {
	for {
		select {
		case s := <-message:
			go buildAlbum(s)
		default:
			time.Sleep(time.Second * 20)
			fmt.Println("default")
		}
	}
}

func buildAlbum(album string) {
	albumHelper := NewAlbumHelper()
	albumHelper.GetAlbum(album)
}
