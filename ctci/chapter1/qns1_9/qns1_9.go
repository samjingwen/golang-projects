package qns1_9

import "strings"

func StringRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	joined := append([]rune(s1), []rune(s1)...)
	return strings.Contains(string(joined), s2)
}
