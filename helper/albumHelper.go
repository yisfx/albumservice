package helper

import (
	"encoding/json"
	"path"
	"strings"

	"../model"
)

const AMBUM_JSON = "album.json"

type AlbumHelper struct {
}

func (this *AlbumHelper) BuildAlbumList(dirPath string) []model.Album {
	pathList := GetFloderListFromPath(dirPath)
	albumList := []model.Album{}
	if pathList == nil {
		return albumList
	}
	for _, album := range pathList {
		albumConfPath := path.Join(dirPath, album, AMBUM_JSON)
		albumConf := &model.Album{}
		confStr := GetFileContentByName(path.Join(albumConfPath))
		json.Unmarshal([]byte(confStr), albumConf)
		if albumConf != nil {
			albumConf.Path = path.Join(dirPath, album)
			///albumConf.PicList = BuildPicForAlbum(*albumConf)
			albumList = append(albumList, *albumConf)
		}
	}
	return albumList
}

func (this *AlbumHelper) GetAlbum(dirPath string) model.Album {
	albumConfPath := path.Join(dirPath, AMBUM_JSON)
	albumConf := &model.Album{}
	confStr := GetFileContentByName(path.Join(albumConfPath))
	json.Unmarshal([]byte(confStr), albumConf)
	if albumConf != nil {
		albumConf.Path = dirPath
		albumConf.PicList = BuildPicForAlbum(*albumConf)
	}
	return *albumConf
}

func getPicName(picName string) string {
	return strings.Split(picName, "-")[0]
}

func BuildPicForAlbum(album model.Album) []model.Picture {
	fileList := GetFileListByPath(album.Path)
	picList := []model.Picture{}
	if fileList == nil {
		return picList
	}
	p := make(map[string]int)
	for _, pic := range fileList {
		if IsPic(pic) {
			name := getPicName(pic)
			if _, ok := p[name]; !ok {
				p[name] = 1
			}
		}
	}
	for n := range p {
		picList = append(picList, model.Picture{
			Name:     n,
			MiniPath: path.Join(album.Path, n+"-mini.jpg"),
			MaxPath:  path.Join(album.Path, n+"-max.jpg"),
			OrgPath:  path.Join(album.Path, n+"-org.jpg"),
			Album:    album.Name,
		})
	}
	return picList
}

func (this *AlbumHelper) ExistsAlbum(albumName string, path string) bool {
	albumList := this.BuildAlbumList(path)
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
	CreateFolder(album.Path)
	///write AMBUM_JSON
	content, _ := json.Marshal(album)
	WriteFile(string(content), path.Join(album.Path, AMBUM_JSON))
}
func (this *AlbumHelper) EditAlbum(album model.Album) {
	content, _ := json.Marshal(album)
	WriteFile(string(content), path.Join(album.Path, AMBUM_JSON))
}

func NewAlbumHelper() *AlbumHelper {
	return &AlbumHelper{}
}
