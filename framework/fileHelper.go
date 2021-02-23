package framework

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetFloderListFromPath(path string) []string {
	fildorList, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}
	list := []string{}
	for _, folder := range fildorList {
		if folder.IsDir() {
			list = append(list, folder.Name())
		}
	}
	return list
}

func GetFileListByPath(path string) []string {
	fileList, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}
	list := []string{}
	for _, file := range fileList {
		if !file.IsDir() {
			list = append(list, file.Name())
		}
	}
	return list
}

func IsPic(fileName string) bool {
	name := strings.Split(fileName, ".")
	if len(name) == 2 && strings.EqualFold(name[1], "jpg") {
		return true
	}
	return false
}

func GetFileContentByName(name string) string {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		return ""
	}
	return string(f)
}

func CreateFolder(dirPath string) {
	os.Mkdir(dirPath, os.ModePerm)
}

func WriteFile(content string, fileName string) {

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 6)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(content)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func DeleteFile(path string) bool {
	if FileExists(path) {
		os.Remove(path)
	}
	return true
}
