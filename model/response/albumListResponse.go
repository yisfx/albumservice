package response

import model ".."

type AlbumListResponse struct {
	BaseResponse
	AlbumList []model.Album
}
