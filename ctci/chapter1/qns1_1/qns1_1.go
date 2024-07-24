package qns1_1

import (
	"sort"
	"strings"
	"unicode/utf8"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func IsUnique1(str string) bool {
	var cache [128]bool
	for _, char := range str {
		val := char - 'a'
		if cache[val] {
			return false
		}
		cache[val] = true
	}
	return true
}

func IsUnique2(str string) bool {
	s := sortString(str)
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return false
		}
	}
	return true
}

func IsUnique3(str string) bool {
	checker := 0
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		val := chars[i] - 'a'
		if checker&(1<<val) > 0 {
			return false
		}
		checker |= 1 << val
	}
	return true
}

// https://www.reddit.com/r/golang/comments/tawo0e/started_learning_go_recently/
func IsUnique4(str string) bool {
	const rRuneError = utf8.RuneError
	const sRuneError = string(rRuneError)
	unique := make(map[rune]struct{})
	for i, r := range str {
		if r == rRuneError {
			if !strings.HasPrefix(str[i:], sRuneError) {
				// ignore UTF-8 encoding errors
				continue
			}
		}
		if _, ok := unique[r]; ok {
			return false
		}
		unique[r] = struct{}{}
	}
	return true
}
