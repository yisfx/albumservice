package albumtool

import (
	"time"

	"../framework"
	model "../model"
)

type PictureProcesser struct {
	sysConf model.SysConf
}

var message = make(chan string)

func In(album string) {
	message <- album
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
	album := albumHelper.GetAlbum(albumPath)
	for _, pic := range album.PicList {
		if !framework.FileExists(pic.MaxPath) {
			buildMaxPic(pic.OrgPath, pic.MaxPath)
		}
		if !framework.FileExists(pic.MiniPath) {
			buildMiniPic(pic.MaxPath, pic.MiniPath)
		}
	}
}

func buildMaxPic(org string, max string) {

}
func buildMiniPic(max string, mini string) {

}
