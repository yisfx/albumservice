package model

import "albumservice/framework/utils"

///相册
type Album struct {
	Name        string
	CNName      string
	Cover       string
	Date        string
	Path        string
	PicList     []*Picture
	Description string
}

///图片
type Picture struct {
	Name     string
	MiniPath string
	MaxPath  string
	OrgPath  string
	Album    string
}

const (
	DeleteImage        = "Image"
	DeleteAbbreviation = "Abbreviation"
)

type PictureUri struct {
	AlbumName string `json:"a"`
	Name      string `json:"n"`
	Type      string `json:"t"`
	Datetime  string `json:"d"`
}

type AlbumList []*Album

func (I AlbumList) Len() int {
	return len(I)
}
func (I AlbumList) Less(i, j int) bool {

	dateI := utils.DateTime.Parse(I[i].Date)
	dateJ := utils.DateTime.Parse(I[j].Date)
	if dateI.Year != dateJ.Year {
		return dateI.Year > dateJ.Year
	}
	if dateI.Month != dateJ.Month {
		return dateI.Month > dateJ.Month
	}
	if dateI.Day != dateJ.Day {
		return dateI.Day > dateJ.Day
	}
	if dateI.Hour != dateJ.Hour {
		return dateI.Hour > dateJ.Hour
	}
	return false
}
func (I AlbumList) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}
