package response

import (
	model "../"
)

type GetAlbumPicListResponse struct {
	BaseResponse
	Album model.Album
}
