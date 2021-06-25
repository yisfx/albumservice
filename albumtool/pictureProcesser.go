package albumtool

import (
	"albumservice/albumtool/albumUtils"
	"albumservice/framework/fileTool"
	"albumservice/framework/model"
	"albumservice/framework/utils"
	albumModel "albumservice/model"
	"albumservice/model/albumConst"
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
	defer utils.ErrorHandler()
	albumHelper := NewAlbumHelper()
	album := albumHelper.GetAlbum(albumName)

	for _, pic := range album.PicList {
		BuildPicture(&pic)
	}
}

func BuildPicture(pic *albumModel.Picture) {
	defer utils.ErrorHandler()
	if !fileTool.FileExists(pic.MaxPath) {
		albumUtils.CompressPicture(pic.OrgPath, pic.MaxPath, albumConst.MaxExtension)
	}
	if !fileTool.FileExists(pic.MiniPath) {
		albumUtils.CompressPicture(pic.MaxPath, pic.MiniPath, albumConst.MiniExtension)
	}
}
