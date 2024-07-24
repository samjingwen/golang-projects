package qns8_7

// PermutationsWithoutDups Write a method to compute
// all permutations of a string of unique characters.
func PermutationsWithoutDups(str string) []string {
	var res []string

	var iter func(curr, prefix []rune)
	iter = func(curr, prefix []rune) {
		if len(curr) == 0 {
			res = append(res, string(prefix))
			return
		}

		for i, ch := range curr {
			var rem []rune
			rem = append(rem, curr[0:i]...)
			rem = append(rem, curr[i+1:]...)

			iter(rem, append(prefix, ch))
		}
	}

	iter([]rune(str), []rune(""))
	return res
}
