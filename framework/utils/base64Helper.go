package utils

import (
	"encoding/base64"
	"os"
)

func Base64ToImage(base64Str string, targetFile string) {
	defer ErrorHandler()
	dist, _ := base64.StdEncoding.DecodeString(base64Str)
	//写入新文件
	f, _ := os.OpenFile(targetFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
}
