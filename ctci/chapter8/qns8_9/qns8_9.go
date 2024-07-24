package qns8_9

// Parens Implement an algorithm to print all valid (e.g., properly opened and closed)
// combinations of n pairs of parentheses.
func Parens(n int) map[string]struct{} {
	if n == 1 {
		return map[string]struct{}{"()": {}}
	}

	res := make(map[string]struct{})
	parens := Parens(n - 1)
	for paren := range parens {
		res[paren+"()"] = struct{}{}
		res["()"+paren] = struct{}{}
		res["("+paren+")"] = struct{}{}
	}
	return res
}

func Parens2(n int) map[string]struct{} {
	res := make(map[string]struct{})

	var iter func(leftCount, rightCount, index int, prefix []rune)
	iter = func(leftCount, rightCount, index int, prefix []rune) {
		if leftCount < 0 || rightCount < leftCount {
			// invalid state
			return
		}
		if leftCount == 0 && rightCount == 0 {
			res[string(prefix)] = struct{}{}
			return
		}

		prefix[index] = '('
		iter(leftCount-1, rightCount, index+1, prefix)
		prefix[index] = ')'
		iter(leftCount, rightCount-1, index+1, prefix)
	}

	iter(n, n, 0, make([]rune, n*2))
	return res
}
