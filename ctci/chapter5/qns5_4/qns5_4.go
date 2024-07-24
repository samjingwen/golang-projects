package qns5_4

func NextBiggest(num uint32) uint32 {
	count0s := 0
	count1s := 0

	val := num
	for val&1 == 0 && val != 0 { // count trailing 0s
		count0s++
		val >>= 1
	}

	for val&1 == 1 { // count 1s after trailing 0s
		count1s++
		val >>= 1
	}

	idx := count0s + count1s

	num = updateBit(num, uint32(idx), true)

	num = num & (^uint32(0) << idx) // clear bit after position idx
	num = num | ((1 << (count1s - 1)) - 1)

	return num
}

func NextSmallest(num uint32) uint32 {
	count0s := 0
	count1s := 0

	val := num
	for val&1 == 1 { // count trailing 1s
		count1s++
		val >>= 1
	}

	for val&1 == 0 && val != 0 {
		count0s++
		val >>= 1
	}

	idx := count0s + count1s

	// clear bit at position idx
	num = updateBit(num, uint32(idx), false)

	// clear bits right of position idx
	num = num & (^uint32(0) << idx)

	// create sequence of 1s and 0s
	var mask uint32 = ((1 << (count1s + 1)) - 1) << (count0s - 1)

	// set count1s number of bits left of position idx
	num = num | mask

	return num
}

func getBit(num uint32, i uint32) bool {
	return (num & (1 << i)) != 0
}

// clear bit at i
func clearBit(num uint32, i uint32) uint32 {
	var mask uint32 = ^(1 << i)
	return num & mask
}

func updateBit(num uint32, i uint32, bitIs1 bool) uint32 {
	var value uint32 = 0
	if bitIs1 {
		value = 1
	}

	var mask uint32 = ^(1 << i)
	return (num & mask) | (value << i)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
