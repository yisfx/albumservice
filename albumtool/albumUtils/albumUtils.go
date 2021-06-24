package albumUtils

import (
	"albumservice/framework/utils"
	"albumservice/model"
	"albumservice/model/albumConst"
	"image"
	"os"
	"path"
	"strings"
)

func BuildPictureModel(album *model.Album, pictureName string) *model.Picture {
	pic := model.Picture{
		Name:     pictureName,
		MiniPath: path.Join(album.Path, pictureName+albumConst.MiniExtension),
		MaxPath:  path.Join(album.Path, pictureName+albumConst.MaxExtension),
		OrgPath:  path.Join(album.Path, pictureName+albumConst.OrgExtension),
		Album:    album.Name,
	}
	return &pic
}

func GetPicName(picName string) string {
	picName = strings.ToLower(picName)
	picName = strings.ReplaceAll(picName, albumConst.MiniExtension, "")
	picName = strings.ReplaceAll(picName, albumConst.MaxExtension, "")
	picName = strings.ReplaceAll(picName, albumConst.OrgExtension, "")
	names := strings.Split(picName, ".")
	return names[0]
}

func CompressPicture(orgPath string, targetPath string, picType string) {
	defer utils.ErrorHandler()
	file, _ := os.Open(orgPath)
	c, _, _ := image.DecodeConfig(file)
	max := 0
	if c.Width > c.Height {
		max = c.Width
	} else {
		max = c.Height
	}
	file.Close()
	if picType == albumConst.MiniExtension {
		max = max / 2
	}
	w := uint(max)
	utils.CompressImg(orgPath, w, targetPath)
}
