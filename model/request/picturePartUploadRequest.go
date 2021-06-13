package request

type PicturePartUploadRequest struct {
	PartIndex   int
	Value       string
	PictureName string
	AlbumName   string
	IsLastPart  bool
}
