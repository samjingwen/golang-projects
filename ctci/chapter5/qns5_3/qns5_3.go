package qns5_3

func FlipBitToWin(num int32) int {
	var arr []int

	count := 0

	var curr int32 = 0

	for i := 0; i < 32; i++ {
		if num&1 != curr {
			arr = append(arr, count)

			// flip 1 to 0, 0 to 1
			curr ^= 1
			count = 0
		}
		count++

		num >>= 1
	}

	arr = append(arr, count)

	res := 1

	for i := 0; i < len(arr); i += 2 {
		zeroes := arr[i]
		leftOnes := 0
		if i-1 > 0 {
			leftOnes = arr[i-1]
		}
		rightOnes := 0
		if i+1 < len(arr) {
			rightOnes = arr[i+1]
		}

		currRes := 0
		if zeroes == 1 {
			currRes = leftOnes + rightOnes + 1
		} else if zeroes > 1 {
			currRes = max(leftOnes, rightOnes) + 1
		} else {
			currRes = max(leftOnes, rightOnes)
		}
		res = max(res, currRes)
	}
	return res
}

func FlipBitToWin2(num int32) int {
	if num == -1 {
		return 32
	}

	currLen := 0
	prevLen := 0
	maxLen := 1
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			// current bit is 1
			currLen++
		} else {
			// current bit is 0
			if num&2 == 0 {
				// next bit is 0
				prevLen = 0
			} else {
				//next bit is 1
				prevLen = currLen
			}
			currLen = 0
		}
		maxLen = max(maxLen, prevLen+currLen+1)
		num >>= 1
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
