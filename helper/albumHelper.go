package helper

import (
	"encoding/json"
	"path"
	"strings"

	"../model"
)

type AlbumHelper struct {
}

func (this *AlbumHelper) BuildAlbumList(dirPath string) []model.Album {
	pathList := GetFloderListFromPath(dirPath)
	albumList := []model.Album{}
	if pathList == nil {
		return albumList
	}
	for _, album := range pathList {
		albumConfPath := path.Join(dirPath, album, "album.json")
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
