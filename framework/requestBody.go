package framework

import (
	"fmt"
	"io"
	"io/ioutil"
)

func ReadBody(body io.Reader) []byte {
	b, _ := ioutil.ReadAll(body)
	fmt.Println(string(b))
	return b
}
