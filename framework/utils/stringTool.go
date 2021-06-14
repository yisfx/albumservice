package utils

import "strings"

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
