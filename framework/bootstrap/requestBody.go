package bootstrap

import (
	"albumservice/framework/utils"
	"io"
	"io/ioutil"
)

func ReadBody(body io.Reader) []byte {
	defer utils.ErrorHandler()
	b, _ := ioutil.ReadAll(body)
	return b
}
