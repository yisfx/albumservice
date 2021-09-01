package utils

import (
	"encoding/json"
	"strings"
)

func IsExist(array []string, find string, ingoreCase bool) bool {
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

func SerializerToJson(i interface{}) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
