
package request

import (
	"albumservice/model"
)

type AddWordRequest struct{
	Title string `json:"Title"`
	Word []*model.Word `json:"Word"`
}