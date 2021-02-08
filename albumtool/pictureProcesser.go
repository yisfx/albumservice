package albumtool

import (
	"time"

	model "../model"
)

type PictureProcesser struct {
	sysConf model.SysConf
}

var message = make(chan string)

func In(album string) {

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
			time.Sleep(time.Second * 2)
		}
	}
}

func buildAlbum(albumPath string) {
	albumHelper := NewAlbumHelper()
	albumHelper.GetAlbum(albumPath)
}
