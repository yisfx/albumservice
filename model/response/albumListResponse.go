package response

import (
	"albumservice/framework/bootstrapmodel"
	model "albumservice/model"
)

type AlbumListResponse struct {
	bootstrapmodel.BaseResponse
	AlbumList []*model.Album
}

type GetAlbumPicListResponse struct {
	bootstrapmodel.BaseResponse
	Album model.Album
}
