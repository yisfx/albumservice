package albumtool

import (
	"image"
	"os"
	"time"

	"albumservice/framework"
	model "albumservice/model"
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

		file, _ := os.Open(pic.OrgPath)
		c, _, _ := image.DecodeConfig(file)
		max := 0
		if c.Width > c.Height {
			max = c.Width
		} else {
			max = c.Height
		}
		file.Close()
		if !framework.FileExists(pic.MaxPath) {
			buildMaxPic(pic.OrgPath, pic.MaxPath, max)
		}
		if !framework.FileExists(pic.MiniPath) {
			buildMiniPic(pic.MaxPath, pic.MiniPath, max)
		}
	}
}

func buildMaxPic(org string, max string, width int) {
	w := uint(width)
	framework.CompressImg(org, w, max)
}
func buildMiniPic(max string, mini string, width int) {
	w := uint(width / 2)
	framework.CompressImg(max, w, mini)
}
