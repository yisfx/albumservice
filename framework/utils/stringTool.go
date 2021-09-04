package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

var StringTool *stringTool

func init() {
	StringTool = &stringTool{}
}

type stringTool struct {
}

func (tool *stringTool) IsExist(array []string, find string, ingoreCase bool) bool {
	for _, a := range array {
		if a == find {
			return true
		}
		if ingoreCase && strings.EqualFold(a, find) {
			return true
		}
	}
	return false
}

func (tool *stringTool) SerializerToJson(i interface{}) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (tool *stringTool) DeSerializerFromJson(str string, v interface{}) interface{} {
	if err := json.Unmarshal([]byte(str), v); err != nil {
		fmt.Println("err:", err)
		return nil
	}
	return v
}
