package response

import (
	"albumservice/model"
)

type GetWordResponse struct{
	Word []*model.Word
}