package qns1_5

func OneAway1(shorter, longer string) bool {
	if len(shorter) > len(longer) {
		return OneAway1(longer, shorter)
	}
	if len(longer)-len(shorter) > 1 {
		return false
	}
	map1 := getCharCount(shorter)
	map2 := getCharCount(longer)

	edited := 0
	for char2, count2 := range map2 {
		count1 := map1[char2]
		edited += abs(count2 - count1)
		if edited > 1 {
			return false
		}
	}
	return true
}

func OneAway2(str1, str2 string) bool {
	if len(str1) > len(str2) {
		return OneAway2(str2, str1)
	}
	if len(str2)-len(str1) > 1 {
		return false
	}

	len1 := len(str1)
	len2 := len(str2)

	i, j := 0, 0
	foundDifference := false
	for i < len1 && j < len2 {
		if str1[i] != str2[j] {
			if foundDifference {
				return false
			}
			foundDifference = true
			if len1 == len2 {
				i++
			}
		} else {
			i++
		}
		j++
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getCharCount(str string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range str {
		charCount[char] += 1
	}
	return charCount
}
