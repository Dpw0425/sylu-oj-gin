package utils

import "strings"

func ArrToString(arr []string) string {
	result := strings.Join(arr, ",")

	return result
}
