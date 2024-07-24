package qns8_1

func TripleStep(n int) int {
	cache := make([]int, n+1)
	return tripleStep(n, cache)
}

func tripleStep(n int, cache []int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 4
	}
	if cache[n] == 0 {
		cache[n] = tripleStep(n-1, cache) + tripleStep(n-2, cache) + tripleStep(n-3, cache)
	}
	return cache[n]
}
