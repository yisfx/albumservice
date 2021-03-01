package framework

import (
	"io"
	"io/ioutil"
)

func ReadBody(body io.Reader) []byte {
	b, _ := ioutil.ReadAll(body)
	return b
}
