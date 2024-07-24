package qns1_4

import (
	"fmt"
	"strings"
	"unicode"
)

func PalindromePermutation1(str string) bool {
	charCountMap := getCharCountMap(str)
	checker := 0
	for _, count := range charCountMap {
		if count%2 != 0 {
			checker += 1
			if checker > 1 {
				return false
			}
		}
	}
	return true
}

func PalindromePermutation2(str string) bool {
	bitVector := createBitVector(strings.ToLower(str))
	return bitVector == 0 || checkExactlyOneBitSet(bitVector)
}

func createBitVector(str string) int32 {
	var bitVector int32 = 0
	for _, char := range str {
		bitVector = toggle(bitVector, getCharNumber(char))
	}
	return bitVector
}

func toggle(bitVector int32, char int32) int32 {
	if char < 0 {
		return bitVector
	}
	var mask int32 = 1 << char
	if bitVector&mask == 0 {
		// Odd number of times character appeared
		// set char bit to 1
		bitVector |= mask
	} else {
		// Even number of times character appeared
		// set char bit to 0
		bitVector &= ^mask
	}
	return bitVector
}

// 00010000 - 1 = 00001111
// 00010000 & 00001111 = 0
func checkExactlyOneBitSet(bitVector int32) bool {
	return (bitVector & (bitVector - 1)) == 0
}

func getCharNumber(char int32) int32 {
	a := 'a'
	z := 'z'
	fmt.Println(">>>", char)
	if a <= char && char <= z {
		return char - a
	}
	return -1
}

func getCharCountMap(str string) map[rune]int {
	charCount := make(map[rune]int)
	str = strings.ToLower(str)
	for _, char := range str {
		if unicode.IsLetter(char) {
			_, exists := charCount[char]
			if exists {
				charCount[char] += 1
			} else {
				charCount[char] = 1
			}
		}
	}
	return charCount
}
