package response

import (
	"albumservice/framework/bootstrapmodel"
	"albumservice/model"
)

type GetAllYearsResponse struct {
	bootstrapmodel.BaseResponse
	AllYears []YearAlbumList
}

type YearAlbumList struct {
	Year      string
	AlbumList []model.Album
}
