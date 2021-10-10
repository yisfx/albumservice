package response

import (
	"albumservice/framework/bootstrapmodel"
)

type DeEntryImageResponse struct {
	bootstrapmodel.BaseResponse
	ImagePath string
}
