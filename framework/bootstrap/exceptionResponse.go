package bootstrap

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Response404(resp http.ResponseWriter, httpMethodName string, request *http.Request) {
	fmt.Println("404 :", httpMethodName, request.URL.Path)
	resss, _ := json.Marshal(fmt.Sprint("404:     ", httpMethodName, ":", request.URL.Path))
	resp.Write(resss)
	return
}

func Response415(resp http.ResponseWriter, httpMethodName string, request *http.Request) {
	fmt.Println("415 :", httpMethodName, request.URL.Path)
	resss, _ := json.Marshal(fmt.Sprint("415:     ", httpMethodName, ":", request.URL.Path))
	resp.Write(resss)
	return
}
