package albumtool

import (
	"albumservice/albumtool/albumUtils"
	"albumservice/framework/fileTool"
	"albumservice/framework/redisTool"
	"albumservice/framework/utils"
	"albumservice/model"
	"albumservice/model/albumConst"
	"encoding/json"
	"fmt"
	"path"
	"strings"
)

const AMBUM_JSON = "album.json"
const (
	Album_List_Key         = "album_list"          //list
	Album_Name_Key         = "album_"              //string  album_aaa
	Album_Picture_List_Key = "album_picture_list_" //list album_picture_list_aaa
	//TODO:
	Picture_Key = "picture_" //string picture_IMG_20210505_115601

	Picture_Cache_Key = "picture_Cache_" // string picture_Cache_aaa_IMG_20210505_115601_index
)

type AlbumHelper struct {
}

func (albumHelper *AlbumHelper) GetAlbumList() []model.Album {
	defer utils.ErrorHandler()
	albumNameList := redisTool.GetList(Album_List_Key)
	albumList := []model.Album{}
	for _, albumName := range albumNameList {
		confStr := redisTool.GetString(Album_Name_Key + albumName)
		if confStr == "" {
			return nil
		}
		albumConf := &model.Album{}
		json.Unmarshal([]byte(confStr), albumConf)
		if albumConf != nil {
			albumList = append(albumList, *albumConf)
		}
	}
	return albumList
}

func (albumHelper *AlbumHelper) GetAlbum(albumName string) *model.Album {
	defer utils.ErrorHandler()
	confStr := redisTool.GetString(Album_Name_Key + albumName)
	if confStr == "" {
		return nil
	}
	albumConf := &model.Album{}
	json.Unmarshal([]byte(confStr), albumConf)
	if albumConf != nil {

		albumConf.PicList = albumHelper.GetPicForAlbum(albumName)
	}
	return albumConf
}

func (thialbumHelpers *AlbumHelper) GetPicForAlbum(albumName string) []model.Picture {
	defer utils.ErrorHandler()
	var picList []model.Picture
	picList = []model.Picture{}

	picNameList := redisTool.GetList(Album_Picture_List_Key + albumName)
	for _, picName := range picNameList {
		pic := &model.Picture{}
		str := redisTool.GetString(Picture_Key + picName)
		json.Unmarshal([]byte(str), pic)
		if pic != nil {
			picList = append(picList, *pic)
		}
	}
	return picList
}

func (albumHelper *AlbumHelper) BuildPicForAlbum(album *model.Album) {
	defer utils.ErrorHandler()
	fileList := fileTool.GetFileListByPath(album.Path)
	if fileList == nil {
		return
	}
	p := make(map[string]int)
	for _, pic := range fileList {
		if fileTool.IsPic(pic) {
			name := albumUtils.GetPicName(pic)
			if _, ok := p[name]; !ok {
				p[name] = 1
			}
		}
	}

	picList := redisTool.GetList(Album_Picture_List_Key + album.Name)
	for n := range p {
		pic := albumUtils.BuildPictureModel(album, n)
		picData, err := json.Marshal(pic)
		if err == nil {
			if !utils.IsExist(picList, n, true) {
				redisTool.SetList(Album_Picture_List_Key+album.Name, n)
			}
			redisTool.SetString(Picture_Key+n, string(picData))
		}
	}
}
func (albumHelper *AlbumHelper) BuildAlbumList(dirPath string) {
	defer utils.ErrorHandler()
	pathList := fileTool.GetFloderListFromPath(dirPath)
	if pathList == nil {
		return
	}
	for _, album := range pathList {

		albumConfPath := path.Join(dirPath, album, AMBUM_JSON)
		albumConf := &model.Album{}
		confStr := fileTool.GetFileContentByName(path.Join(albumConfPath))
		if confStr == "" {
			continue
		}
		json.Unmarshal([]byte(confStr), albumConf)

		albumConf.Path = path.Join(dirPath, album)

		if !albumHelper.ExistsAlbum(album) {
			///save albumList
			redisTool.SetList(Album_List_Key, albumConf.Name)
		}

		///save album conf
		value, err := json.Marshal(albumConf)
		if err == nil {
			redisTool.SetString(Album_Name_Key+albumConf.Name, string(value))
		}
	}
}

func (albumHelper *AlbumHelper) ExistsAlbum(albumName string) bool {
	defer utils.ErrorHandler()
	albumList := albumHelper.GetAlbumList()
	b := false
	for _, a := range albumList {
		if strings.EqualFold(a.Name, albumName) {
			b = true
			break
		}
	}
	return b
}

func (albumHelper *AlbumHelper) CreateAlbum(album model.Album) {
	defer utils.ErrorHandler()
	///create folder
	fileTool.CreateFolder(album.Path)
	///write AMBUM_JSON
	content, _ := json.Marshal(album)
	fileTool.WriteFile(string(content), path.Join(album.Path, AMBUM_JSON))

	redisTool.SetList(Album_List_Key, album.Name)
	redisTool.SetString(Album_Name_Key+album.Name, string(content))
}

func (albumHelper *AlbumHelper) EditAlbum(album model.Album) {
	defer utils.ErrorHandler()
	content, _ := json.Marshal(album)

	fileTool.WriteFile(string(content), path.Join(album.Path, AMBUM_JSON))
	redisTool.SetString(Album_Name_Key+album.Name, string(content))
}

func (albumHelper *AlbumHelper) AddAlbumPicture(album *model.Album, pictureName string) {
	defer utils.ErrorHandler()
	pic := albumUtils.BuildPictureModel(album, pictureName)
	picData, err := json.Marshal(pic)
	if err == nil {
		redisTool.SetList(Album_Picture_List_Key+album.Name, pictureName)
		redisTool.SetString(Picture_Key+pictureName, string(picData))
	}
}

func (albumHelper *AlbumHelper) DeleteAlbumPic(album *model.Album, picName string, deleteType string) {
	defer utils.ErrorHandler()
	///org
	if deleteType == model.DeleteImage {
		fileTool.DeleteFile(album.Path + "/" + picName + albumConst.OrgExtension)
		fileTool.DeleteFile(album.Path + "/" + picName + albumConst.MaxExtension)
		fileTool.DeleteFile(album.Path + "/" + picName + albumConst.MiniExtension)
		redisTool.DeleteList(Album_Picture_List_Key+album.Name, picName)
		redisTool.DelKey(Picture_Key + picName)
	}
	///max
	if deleteType == model.DeleteAbbreviation {
		//max
		fileTool.DeleteFile(album.Path + "/" + picName + albumConst.MaxExtension)
		//mini
		fileTool.DeleteFile(album.Path + "/" + picName + albumConst.MiniExtension)
	}
}

func (albumHelper *AlbumHelper) CacheUploadImage(albumName string, pictureName string, index int, cacheData string) {
	defer utils.ErrorHandler()
	redisTool.SetTempCache(fmt.Sprint(Picture_Cache_Key, albumName, "_", pictureName, "_", index), cacheData)
}

func (albumHelper *AlbumHelper) BuildCacheUploadImage(albumName string, pictureName string, lastIndex int) {

	defer utils.ErrorHandler()

	cacheData := *new([]string)
	cacheKey := fmt.Sprint(Picture_Cache_Key, albumName, "_", pictureName, "_")
	for index := 0; index <= lastIndex; index++ {
		str := redisTool.GetString(fmt.Sprint(cacheKey, index))
		if str != "" {
			cacheData = append(cacheData, str)
			redisTool.DelKey(fmt.Sprint(cacheKey, index))
		}
	}
	album := albumHelper.GetAlbum(albumName)
	pictureName = album.Name + "-" + pictureName
	///save base64 image
	orgPath := path.Join(album.Path, pictureName+".jpg")
	utils.Base64ToImage(strings.Join(cacheData, ""), orgPath)

	orgPicture := path.Join(album.Path, pictureName+albumConst.OrgExtension)
	///save org picture & compress
	albumUtils.CompressPicture(orgPath, orgPicture, albumConst.OrgExtension)
	albumHelper.AddAlbumPicture(album, pictureName)

	fileTool.DeleteFile(orgPath)
}

func NewAlbumHelper() *AlbumHelper {
	return &AlbumHelper{}
}
