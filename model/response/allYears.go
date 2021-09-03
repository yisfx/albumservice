package response

import "albumservice/framework/bootstrapmodel"

type GetAllYearsResponse struct {
	bootstrapmodel.BaseResponse
	AllYears []string
}
