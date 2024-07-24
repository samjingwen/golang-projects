package qns1_2

// CheckPermutation Given two strings, write a method to
// decide if one is a permutation of the other.
func CheckPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1CharCountMap := getCharCountMap(s1)
	s2CharCountMap := getCharCountMap(s2)

	for k, s1Count := range s1CharCountMap {
		s2Count, exists := s2CharCountMap[k]
		if s1Count > s2Count {
			return false
		} else if !exists {
			return false
		}
	}
	return true
}

func getCharCountMap(str string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range str {
		_, exists := charCount[char]
		if exists {
			charCount[char] += 1
		} else {
			charCount[char] = 1
		}
	}
	return charCount
}
