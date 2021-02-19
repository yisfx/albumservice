package albumtool

import (
	"fmt"
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
		file.Close()
		fmt.Println("width = ", c.Width, c.Height)

		if !framework.FileExists(pic.MaxPath) {
			buildMaxPic(pic.OrgPath, pic.MaxPath)
		}
		if !framework.FileExists(pic.MiniPath) {
			buildMiniPic(pic.MaxPath, pic.MiniPath)
		}
	}
}

func buildMaxPic(org string, max string) {
	// fmt.Println(org, max)
	// framework.CompressImg(org, 2560, max)
}
func buildMiniPic(max string, mini string) {
	// fmt.Println(max, mini)
	// framework.CompressImg(max, 500, mini)
}
