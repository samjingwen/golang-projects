package qns8_5

func RecursiveMultiply(a, b int) int {
	cache := make(map[int]int)

	var iter func(a, b int) int
	iter = func(a, b int) int {
		if a < b {
			return iter(b, a)
		}
		if b == 0 {
			return 0
		}
		if b == 1 {
			return a
		}
		res, ok := cache[b]
		if ok {
			return res
		}

		if b&1 == 1 {
			res = a + iter(a, b>>1) + iter(a, b>>1)
			cache[b] = res
			return res
		}
		res = iter(a, b>>1) + iter(a, b>>1)
		cache[b] = res
		return res
	}

	return iter(a, b)
}
