package albumUtils

import (
	"albumservice/framework/utils"
	"albumservice/model"
	"albumservice/model/albumConst"
	"path"
	"strings"
)

func BuildPictureModel(album *model.Album, pictureName string) *model.Picture {
	defer utils.ErrorHandler()
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
	defer utils.ErrorHandler()
	picName = strings.ToLower(picName)
	picName = strings.ReplaceAll(picName, albumConst.MiniExtension, "")
	picName = strings.ReplaceAll(picName, albumConst.MaxExtension, "")
	picName = strings.ReplaceAll(picName, albumConst.OrgExtension, "")
	names := strings.Split(picName, ".")
	return names[0]
}

func CompressPicture(orgPath string, targetPath string, picType string) {
	quality := 100

	if picType == albumConst.MaxExtension {
		quality = 80
	}
	if picType == albumConst.MiniExtension {
		quality = 50
	}
	utils.CompressJpgResource(orgPath, targetPath, quality)
}
