package albumtool

import (
	"albumservice/albumtool/albumUtils"
	"albumservice/albumtool/constfield"
	"albumservice/framework/fileTool"
	"albumservice/framework/redisTool"
	"albumservice/framework/utils"
	"albumservice/model"
	"albumservice/model/albumConst"
	"encoding/json"
	"fmt"
	"path"
	"sort"
	"strings"
)

type AlbumHelper struct {
}

// GetAlbumList 获取所有AlbumList
func (albumHelper *AlbumHelper) GetAlbumList() []*model.Album {
	albumNameList := redisTool.GetList(constfield.Album_List_Key)
	albumList := []*model.Album{}
	for _, albumName := range albumNameList {
		confStr := redisTool.GetString(constfield.Album_Name_Key + albumName)
		if confStr == "" {
			return nil
		}
		albumConf := &model.Album{}
		json.Unmarshal([]byte(confStr), albumConf)
		if albumConf != nil {
			albumList = append(albumList, albumConf)
		}
	}
	return albumList
}

// GetAlbum 根据AlbumName 获取Album
func (albumHelper *AlbumHelper) GetAlbum(albumName string) *model.Album {
	confStr := redisTool.GetString(constfield.Album_Name_Key + albumName)
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

// GetPicForAlbum 根据AlbumName 获取Album的PicList
func (thialbumHelpers *AlbumHelper) GetPicForAlbum(albumName string) []*model.Picture {
	picList := []*model.Picture{}

	picNameList := redisTool.GetList(constfield.Album_Picture_List_Key + albumName)
	for _, picName := range picNameList {
		pic := &model.Picture{}
		str := redisTool.GetString(constfield.Picture_Key + picName)
		json.Unmarshal([]byte(str), pic)
		if pic != nil {
			picList = append(picList, pic)
		}
	}
	return picList
}

// BuildPicForAlbum 重建Album的Pictrue
func (albumHelper *AlbumHelper) BuildPicForAlbum(album *model.Album) {
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

	picList := redisTool.GetList(constfield.Album_Picture_List_Key + album.Name)
	for n := range p {
		pic := albumUtils.BuildPictureModel(album, n)
		picData, err := json.Marshal(pic)
		if err == nil {
			if !utils.StringTool.IsExist(picList, n, true) {
				redisTool.SetList(constfield.Album_Picture_List_Key+album.Name, n)
			}
			redisTool.SetString(constfield.Picture_Key+n, string(picData))
		}
	}
}

// BuildAlbumList 重建路径下所有的Album
func (albumHelper *AlbumHelper) BuildAlbumList(dirPath string) {
	pathList := fileTool.GetFloderListFromPath(dirPath)
	if pathList == nil {
		return
	}
	for _, album := range pathList {

		albumConfPath := path.Join(dirPath, album, constfield.AMBUM_JSON)
		albumConf := &model.Album{}
		confStr := fileTool.GetFileContentByName(path.Join(albumConfPath))
		if confStr == "" {
			continue
		}
		json.Unmarshal([]byte(confStr), albumConf)

		albumConf.Path = path.Join(dirPath, album)

		if !albumHelper.ExistsAlbum(album) {
			///save albumList
			redisTool.SetList(constfield.Album_List_Key, albumConf.Name)
		}

		///save album conf
		value, err := json.Marshal(albumConf)
		if err == nil {
			redisTool.SetString(constfield.Album_Name_Key+albumConf.Name, string(value))
		}
	}
}

// ExistsAlbum 是否存在AlbumName
func (albumHelper *AlbumHelper) ExistsAlbum(albumName string) bool {
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

// CreateAlbum CreateAlbum
func (albumHelper *AlbumHelper) CreateAlbum(album model.Album) {
	///create folder
	fileTool.CreateFolder(album.Path)
	///write AMBUM_JSON
	content, _ := json.Marshal(album)
	fileTool.WriteFile(string(content), path.Join(album.Path, constfield.AMBUM_JSON))

	redisTool.SetList(constfield.Album_List_Key, album.Name)
	redisTool.SetString(constfield.Album_Name_Key+album.Name, string(content))
}

// EditAlbum EditAlbum
func (albumHelper *AlbumHelper) EditAlbum(album model.Album) {
	content, _ := json.Marshal(album)

	fileTool.WriteFile(string(content), path.Join(album.Path, constfield.AMBUM_JSON))
	redisTool.SetString(constfield.Album_Name_Key+album.Name, string(content))
}

// AddAlbumPicture 往cache加入一个album的图片
func (albumHelper *AlbumHelper) AddAlbumPicture(album *model.Album, pictureName string) {
	pic := albumUtils.BuildPictureModel(album, pictureName)
	picData, err := json.Marshal(pic)
	if err == nil {
		redisTool.SetList(constfield.Album_Picture_List_Key+album.Name, pictureName)
		redisTool.SetString(constfield.Picture_Key+pictureName, string(picData))
	}
}

// DeleteAlbum 删除 一个Album，包括cache * file
func (AlbumHelper *AlbumHelper) DeleteAlbum(album *model.Album) {
	redisTool.DeleteList(constfield.Album_List_Key, album.Name)
	redisTool.DelKey(constfield.Album_Name_Key + album.Name)
	fmt.Println("delete json")
	fileTool.DeleteFile(path.Join(album.Path, constfield.AMBUM_JSON))
	fmt.Println("delete folder")
	fileTool.DeleteFolder(album.Path)
}

// DeleteAlbumPic 根据 deleteType 删除 picture
func (albumHelper *AlbumHelper) DeleteAlbumPic(album *model.Album, picName string, deleteType string) {
	///org
	if deleteType == model.DeleteImage {
		fileTool.DeleteFile(album.Path + "/" + picName + albumConst.OrgExtension)
		fileTool.DeleteFile(album.Path + "/" + picName + albumConst.MaxExtension)
		fileTool.DeleteFile(album.Path + "/" + picName + albumConst.MiniExtension)
		redisTool.DeleteList(constfield.Album_Picture_List_Key+album.Name, picName)
		redisTool.DelKey(constfield.Picture_Key + picName)
	}
	///max
	if deleteType == model.DeleteAbbreviation {
		//max
		fileTool.DeleteFile(album.Path + "/" + picName + albumConst.MaxExtension)
		//mini
		fileTool.DeleteFile(album.Path + "/" + picName + albumConst.MiniExtension)
	}
}

// CacheUploadImage 上传picture片段
func (albumHelper *AlbumHelper) CacheUploadImage(albumName string, pictureName string, index int, cacheData string) {
	redisTool.SetTempCache(fmt.Sprint(constfield.Picture_Cache_Key, albumName, "_", pictureName, "_", index), cacheData)
}

// BuildCacheUploadImage 保存上传的picture
func (albumHelper *AlbumHelper) BuildCacheUploadImage(albumName string, pictureName string, lastIndex int) {
	cacheData := *new([]string)
	cacheKey := fmt.Sprint(constfield.Picture_Cache_Key, albumName, "_", pictureName, "_")
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

// GetAllYears 当前上传的年份
func (albumHelper *AlbumHelper) GetAllYears() []string {
	var rrr = redisTool.GetList(constfield.All_Years)
	return rrr
}

// BuildAllYears buildAllYears
func (albumHeler *AlbumHelper) BuildAllYears() {
	albumList := albumHeler.GetAlbumList()
	yearList := map[string][]*model.Album{}
	for _, album := range albumList {
		date := utils.DateTime.Parse(album.Date)
		yearList[fmt.Sprint(date.Year)] = append(yearList[fmt.Sprint(date.Year)], album)
	}

	redisTool.DelKey(constfield.All_Years)

	for year, al := range yearList {
		redisTool.SetList(constfield.All_Years, year)
		redisTool.DelKey(constfield.Year_Album_List_Key + year)
		for _, a := range al {
			redisTool.SetList(constfield.Year_Album_List_Key+year, a.Name)
		}
	}
}

// GetAlbumListByYear 根据年获取albumList
func (albumHeler *AlbumHelper) GetAlbumListByYear(year string) []*model.Album {
	albumNameList := redisTool.GetList(constfield.Year_Album_List_Key + year)
	result := []*model.Album{}
	for _, albumName := range albumNameList {
		result = append(result, albumHeler.GetAlbum(albumName))
	}
	l := model.AlbumList(result)
	sort.Sort(l)
	return l
}

func NewAlbumHelper() *AlbumHelper {
	return &AlbumHelper{}
}
