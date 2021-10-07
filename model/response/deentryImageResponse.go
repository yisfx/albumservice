package response

import (
	"albumservice/framework/bootstrapmodel"
	model "albumservice/model"
)

type DeEntryImageResponse struct {
	bootstrapmodel.BaseResponse
	Image model.PictureUri
}
