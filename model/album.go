package modal

///相册
type Album struct {
	Name        string
	CNName		string
	Cover       string
	Date        string
	Path        string
	PicList     []Picture
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
