package qns1_6

import (
	"strconv"
)

func StringCompression(str string) string {
	var result []rune

	currentChar := str[0]
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] == currentChar {
			count += 1
		} else {
			result = append(result, rune(currentChar))
			result = append(result, []rune(strconv.Itoa(count))...)
			currentChar = str[i]
			count = 1
		}
	}
	result = append(result, rune(currentChar))
	result = append(result, []rune(strconv.Itoa(count))...)
	if len(result) <= len(str) {
		return string(result)
	}
	return str
}
