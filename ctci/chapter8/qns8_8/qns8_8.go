package qns8_8

// PermutationsWithDups Write a method to compute all permutations of a string whose characters are not necessarily
// unique. The list of permutations should not have duplicates.
func PermutationsWithDups(str string) []string {
	var res []string
	L := len([]rune(str))
	charCount := buildFreqTable(str)

	var iter func(charCount map[rune]int, prefix []rune)
	iter = func(charCount map[rune]int, prefix []rune) {
		if len(prefix) == L {
			res = append(res, string(prefix))
			return
		}

		for ch, count := range charCount {
			if count > 0 {
				charCount[ch] -= 1
				iter(charCount, append(prefix, ch))
				charCount[ch] += 1
			}
		}
	}

	iter(charCount, []rune(""))
	return res
}

func buildFreqTable(str string) map[rune]int {
	cache := make(map[rune]int)
	for _, ch := range str {
		cache[ch] += 1
	}
	return cache
}
