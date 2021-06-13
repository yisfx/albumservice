package albumtool

import (
	framework "albumservice/framework"
	model "albumservice/model"
	"encoding/json"
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
)

type AlbumHelper struct {
}

func (this *AlbumHelper) GetAlbumList() []model.Album {
	var albumNameList []string
	albumNameList = framework.GetList(Album_List_Key)
	albumList := []model.Album{}
	for _, albumName := range albumNameList {
		confStr := framework.GetString(Album_Name_Key + albumName)
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
	confStr := framework.GetString(Album_Name_Key + albumName)
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
	picNameList = framework.GetList(Album_Picture_List_Key + albumName)
	for _, picName := range picNameList {
		pic := &model.Picture{}
		str := framework.GetString(Picture_Key + picName)
		json.Unmarshal([]byte(str), pic)
		if pic != nil {
			picList = append(picList, *pic)
		}
	}
	return picList
}

func (this *AlbumHelper) buildAlbumList(dirPath string) {
	pathList := framework.GetFloderListFromPath(dirPath)
	if pathList == nil {
		return
	}
	for _, album := range pathList {
		albumConfPath := path.Join(dirPath, album, AMBUM_JSON)
		albumConf := &model.Album{}
		confStr := framework.GetFileContentByName(path.Join(albumConfPath))
		json.Unmarshal([]byte(confStr), albumConf)
		if albumConf != nil {
			albumConf.Path = path.Join(dirPath, album)
			///save albumList
			framework.SetList(Album_List_Key, albumConf.Name)
			///save album conf
			value, err := json.Marshal(albumConf)
			if err == nil {
				framework.SetString(Album_Name_Key+albumConf.Name, string(value))
			}
		}
	}
}

func getPicName(picName string) string {

	names := strings.Split(picName, "-")
	if len(names) == 3 {
		return names[0] + "-" + names[1]
	}
	return names[0]
}

func buildPicForAlbum(album model.Album) {
	fileList := framework.GetFileListByPath(album.Path)
	if fileList == nil {
		return
	}
	p := make(map[string]int)
	for _, pic := range fileList {
		if framework.IsPic(pic) {
			name := getPicName(pic)
			if _, ok := p[name]; !ok {
				p[name] = 1
			}
		}
	}
	for n := range p {
		pic := model.Picture{
			Name:     n,
			MiniPath: path.Join(album.Path, n+"-mini.jpg"),
			MaxPath:  path.Join(album.Path, n+"-max.jpg"),
			OrgPath:  path.Join(album.Path, n+"-org.jpg"),
			Album:    album.Name,
		}
		picData, err := json.Marshal(pic)
		if err != nil {
			framework.SetList(Album_Picture_List_Key+album.Name, n)
			framework.SetString(Picture_Key+n, string(picData))
		}
	}
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
	framework.CreateFolder(album.Path)
	///write AMBUM_JSON
	content, _ := json.Marshal(album)
	framework.WriteFile(string(content), path.Join(album.Path, AMBUM_JSON))

	framework.SetList(Album_List_Key, album.Name)
	framework.SetString(Album_Name_Key+album.Name, string(content))
}

func (this *AlbumHelper) EditAlbum(album model.Album) {
	content, _ := json.Marshal(album)
	framework.WriteFile(string(content), path.Join(album.Path, AMBUM_JSON))
	framework.SetString(Album_Name_Key+album.Name, string(content))
}

func (this *AlbumHelper) AddAlbumPicture(albumName string, pictureName string) {
	pic := model.Picture{
		Name:     pictureName,
		MiniPath: path.Join(albumName, pictureName+"-mini.jpg"),
		MaxPath:  path.Join(albumName, pictureName+"-max.jpg"),
		OrgPath:  path.Join(albumName, pictureName+"-org.jpg"),
		Album:    albumName,
	}
	picData, err := json.Marshal(pic)
	if err != nil {
		framework.SetList(Album_Picture_List_Key+albumName, pictureName)
		framework.SetString(Picture_Key+pictureName, string(picData))
	}
}
func (this *AlbumHelper) DeleteAlbumPic(albumPath string, picName string, deleteType string) {
	///org
	if deleteType == model.DeleteImage {
		framework.DeleteFile(albumPath + "/" + picName + "-org.jpg")
		framework.DeleteFile(albumPath + "/" + picName + "-max.jpg")
		framework.DeleteFile(albumPath + "/" + picName + "-mini.jpg")
		framework.DeleteList(Album_Picture_List_Key+picName, picName)
		framework.DelKey(Picture_Key + picName)
	}
	///max
	if deleteType == model.DeleteAbbreviation {
		//max
		framework.DeleteFile(albumPath + "/" + picName + "-max.jpg")
		//mini
		framework.DeleteFile(albumPath + "/" + picName + "-mini.jpg")
	}
}

func NewAlbumHelper() *AlbumHelper {
	return &AlbumHelper{}
}
