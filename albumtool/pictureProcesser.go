package albumtool

import (
	"albumservice/framework/fileTool"
	"albumservice/framework/utils"
	"albumservice/framework/model"
	"image"
	"os"
	"time"
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

func buildAlbum(albumName string) {
	albumHelper := NewAlbumHelper()
	album := albumHelper.GetAlbum(albumName)

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
		if !fileTool.FileExists(pic.MaxPath) {
			buildMaxPic(pic.OrgPath, pic.MaxPath, max)
		}
		if !fileTool.FileExists(pic.MiniPath) {
			buildMiniPic(pic.MaxPath, pic.MiniPath, max)
		}
	}
}

func buildMaxPic(org string, max string, width int) {
	w := uint(width)
	utils.CompressImg(org, w, max)
}
func buildMiniPic(max string, mini string, width int) {
	w := uint(width / 2)
	utils.CompressImg(max, w, mini)
}
