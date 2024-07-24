package qns5_1

func Insertion(N, M, i, j int) int {
	// generate zeroes from j to i
	// if j = 6, i = 2 then 1111 1000 0011
	mask := (((-1 << (j - i + 1)) + 1) << 2) - 1

	// clear bits from j to i
	// if M = 0100 0111 1100, then 0100 0000 0000
	mask = mask & N

	return (M << i) | mask
}
