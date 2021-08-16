package response

import (
	model "albumservice/model"
)

type AlbumListResponse struct {
	BaseResponse
	AlbumList []*model.Album
}
type GetAlbumPicListResponse struct {
	BaseResponse
	Album model.Album
}
