package albumtool

import (
	"albumservice/framework/fileTool"
	"albumservice/framework/redisTool"
	"albumservice/framework/utils"
	model "albumservice/model"
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

func (this *AlbumHelper) GetAlbumList() []model.Album {
	var albumNameList []string
	albumNameList = redisTool.GetList(Album_List_Key)
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

func (this *AlbumHelper) GetAlbum(albumName string) *model.Album {
	confStr := redisTool.GetString(Album_Name_Key + albumName)
	if confStr == "" {
		return nil
	}
	albumConf := &model.Album{}
	json.Unmarshal([]byte(confStr), albumConf)
	if albumConf != nil {

		albumConf.PicList = this.GetPicForAlbum(albumName)
	}
	return albumConf
}

func (this *AlbumHelper) GetPicForAlbum(albumName string) []model.Picture {
	var picList []model.Picture
	picList = []model.Picture{}
	var picNameList []string
	picNameList = redisTool.GetList(Album_Picture_List_Key + albumName)
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
func (this *AlbumHelper) BuildPicForAlbum(album *model.Album) {
	fileList := fileTool.GetFileListByPath(album.Path)
	if fileList == nil {
		return
	}
	p := make(map[string]int)
	for _, pic := range fileList {
		if fileTool.IsPic(pic) {
			name := getPicName(pic)
			if _, ok := p[name]; !ok {
				p[name] = 1
			}
		}
	}

	picList := redisTool.GetList(Album_Picture_List_Key + album.Name)
	for n := range p {
		pic := model.Picture{
			Name:     n,
			MiniPath: path.Join(album.Path, n+"-mini.jpg"),
			MaxPath:  path.Join(album.Path, n+"-max.jpg"),
			OrgPath:  path.Join(album.Path, n+".jpg"),
			Album:    album.Name,
		}
		picData, err := json.Marshal(pic)
		if err == nil {
			if !utils.IsExist(picList, n, true) {
				redisTool.SetList(Album_Picture_List_Key+album.Name, n)
			}
			redisTool.SetString(Picture_Key+n, string(picData))
		}
	}
}
func (this *AlbumHelper) BuildAlbumList(dirPath string) {
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

		if !this.ExistsAlbum(album) {
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

func getPicName(picName string) string {
	picName = strings.ToLower(picName)
	picName = strings.ReplaceAll(picName, "-mini", "")
	picName = strings.ReplaceAll(picName, "-max", "")
	picName = strings.ReplaceAll(picName, "-org", "")
	names := strings.Split(picName, ".")
	return names[0]
}

func (this *AlbumHelper) ExistsAlbum(albumName string) bool {
	albumList := this.GetAlbumList()
	b := false
	for _, a := range albumList {
		if strings.EqualFold(a.Name, albumName) {
			b = true
			break
		}
	}
	return b
}

func (this *AlbumHelper) CreateAlbum(album model.Album) {
	///create folder
	fileTool.CreateFolder(album.Path)
	///write AMBUM_JSON
	content, _ := json.Marshal(album)
	fileTool.WriteFile(string(content), path.Join(album.Path, AMBUM_JSON))

	redisTool.SetList(Album_List_Key, album.Name)
	redisTool.SetString(Album_Name_Key+album.Name, string(content))
}

func (this *AlbumHelper) EditAlbum(album model.Album) {
	content, _ := json.Marshal(album)

	fileTool.WriteFile(string(content), path.Join(album.Path, AMBUM_JSON))
	redisTool.SetString(Album_Name_Key+album.Name, string(content))
}

func (this *AlbumHelper) AddAlbumPicture(album *model.Album, pictureName string) {
	pic := model.Picture{
		Name:     pictureName,
		MiniPath: path.Join(album.Path, pictureName+"-mini.jpg"),
		MaxPath:  path.Join(album.Path, pictureName+"-max.jpg"),
		OrgPath:  path.Join(album.Path, pictureName+"-org.jpg"),
		Album:    album.Name,
	}
	picData, err := json.Marshal(pic)
	if err == nil {
		redisTool.SetList(Album_Picture_List_Key+album.Name, pictureName)
		redisTool.SetString(Picture_Key+pictureName, string(picData))
	}
}
func (this *AlbumHelper) DeleteAlbumPic(album *model.Album, picName string, deleteType string) {
	///org
	if deleteType == model.DeleteImage {
		fileTool.DeleteFile(album.Path + "/" + picName + "-org.jpg")
		fileTool.DeleteFile(album.Path + "/" + picName + "-max.jpg")
		fileTool.DeleteFile(album.Path + "/" + picName + "-mini.jpg")
		redisTool.DeleteList(Album_Picture_List_Key+album.Name, picName)
		redisTool.DelKey(Picture_Key + picName)
	}
	///max
	if deleteType == model.DeleteAbbreviation {
		//max
		fileTool.DeleteFile(album.Path + "/" + picName + "-max.jpg")
		//mini
		fileTool.DeleteFile(album.Path + "/" + picName + "-mini.jpg")
	}
}

func (this *AlbumHelper) CacheUploadImage(albumName string, pictureName string, index int, cacheData string) {
	redisTool.SetTempCache(fmt.Sprint(Picture_Cache_Key, albumName, "_", pictureName, "_", index), cacheData)
}
func (this *AlbumHelper) BuildCacheUploadImage(albumName string, pictureName string, lastIndex int) {
	cacheData := *new([]string)
	cacheKey := fmt.Sprint(Picture_Cache_Key, albumName, "_", pictureName, "_")
	for index := 0; index <= lastIndex; index++ {
		str := redisTool.GetString(fmt.Sprint(cacheKey, index))
		if str != "" {
			cacheData = append(cacheData, str)
			redisTool.DelKey(fmt.Sprint(cacheKey, index))
		}
	}
	album := this.GetAlbum(albumName)
	pictureName = album.Name + "-" + pictureName
	///save base64 image
	utils.Base64ToImage(strings.Join(cacheData, ""), path.Join(album.Path, pictureName+"-org.jpg"))

	this.AddAlbumPicture(album, pictureName)
}

func NewAlbumHelper() *AlbumHelper {
	return &AlbumHelper{}
}
