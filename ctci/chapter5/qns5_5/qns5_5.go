package qns5_5

func Conversion(x, y int) int {
	diff := x ^ y

	count := 0
	for i := 0; i < 32; i++ {
		if diff&1 == 1 {
			count++
		}
		diff >>= 1
	}

	return count
}
