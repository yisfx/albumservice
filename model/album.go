package model

///相册
type Album struct {
	Name  string
	Cover string
	Date  string
	Path  string ///绝对路径
}

const (
	Mini = iota ///缩略图
	Max         ///大图压缩
	Org         ///原图
)

///图片
type Picture struct {
	Name    string
	PicType int ///图片后缀 mini、max、org
	Album   string
}
