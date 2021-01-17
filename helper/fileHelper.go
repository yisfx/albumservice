package helper

import (
	"io/ioutil"
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
