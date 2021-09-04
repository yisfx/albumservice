package albumUtils

import (
	"albumservice/framework/utils"
	"albumservice/model"
	"albumservice/model/albumConst"
	"fmt"
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
	quality := 100

	if picType == albumConst.MaxExtension {
		quality = 50
	}
	if picType == albumConst.MiniExtension {
		quality = 25
	}
	utils.CompressJpgResource(orgPath, targetPath, quality)
}

func EncryptAlbumName(albunName string) string {
	str, _ := utils.DESHelper.DESCBCEncrypt(fmt.Sprintf("%v--%v", albunName, utils.DateTime.Now().ToString()))
	return str
}

func DecryptAlbumName(s string) (string, *utils.Date) {
	str, _ := utils.DESHelper.DESCBCDecrypt(s)
	albumName := strings.Split(str, "--")
	if len(albumName) < 2 {
		return "", &utils.Date{}
	}
	var d = utils.DateTime.Parse(albumName[1])
	if d.IsValid() {
		return albumName[0], d
	}
	return "", nil
}

func EncryptImageUri(albumName, pictureName, ex string) string {
	p := &model.PictureUri{AlbumName: albumName, Name: pictureName, Type: ex, Datetime: utils.DateTime.Now().ToString()}
	uri, err := utils.DESHelper.DESCBCEncrypt(utils.StringTool.SerializerToJson(p))
	if err != nil {
		return ""
	}
	return uri
}

func DecryptImageUri(str string) *model.PictureUri {
	jsonStr, err := utils.DESHelper.DESCBCDecrypt(str)
	if err != nil {
		panic(err)
	}
	p := utils.StringTool.DeSerializerFromJson(jsonStr, &model.PictureUri{})
	ss := p.(*model.PictureUri)
	return ss

}
