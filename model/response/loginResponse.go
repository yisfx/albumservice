package response

import "albumservice/framework/bootstrapmodel"

type LoginResponse struct {
	bootstrapmodel.BaseResponse
	LoginToken string
}
