package qns8_14

func countEval(expr string, result bool) int {

	var eval func(expr []rune, result bool) int
	eval = func(expr []rune, result bool) int {
		if len(expr) == 0 {
			return 0
		}
		if len(expr) == 1 {
			if charToBool(expr[0]) == result {
				return 1
			}
			return 0
		}

		ways := 0
		for i := 1; i < len(expr); i += 2 {
			ch := expr[i]
			left := expr[0:i]
			right := expr[i+1:]

			leftTrue := eval(left, true)
			leftFalse := eval(left, false)
			rightTrue := eval(right, true)
			rightFalse := eval(right, false)

			total := (leftTrue + leftFalse) * (rightTrue + rightFalse)

			totalTrue := 0
			if ch == '^' {
				totalTrue = leftTrue*rightFalse + leftFalse*rightTrue
			} else if ch == '&' {
				totalTrue = leftTrue * rightTrue
			} else if ch == '|' {
				totalTrue = leftTrue*rightTrue + leftFalse*rightTrue + leftTrue*rightFalse
			}

			subWays := 0
			if result {
				subWays = totalTrue
			} else {
				subWays = total - totalTrue
			}
			ways += subWays
		}
		return ways
	}

	return eval([]rune(expr), result)
}

func charToBool(ch rune) bool {
	return ch != '0'
}
