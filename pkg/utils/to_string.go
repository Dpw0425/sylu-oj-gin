package utils

import "strings"

func ArrToString(arr []string) string {
	result := strings.Join(arr, ",")

	return result
}

func StringToArr(str string) []string {
	result := strings.Split(str, ",")

	return result
}
