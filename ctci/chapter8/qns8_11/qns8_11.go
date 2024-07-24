package qns8_11

func Coins(n int) int {
	denoms := []int{25, 10, 5, 1}

	cache := make([][]int, n+1)
	for i, _ := range cache {
		cache[i] = make([]int, len(denoms))
	}

	var iter func(amount, index int) int
	iter = func(amount, index int) int {
		if amount == 0 {
			return 1
		}
		if amount < 0 || index >= len(denoms) {
			return 0
		}
		if cache[amount][index] > 0 {
			return cache[amount][index]
		}

		ways := iter(amount, index+1) + iter(amount-denoms[index], index)
		cache[amount][index] = ways
		return ways
	}

	return iter(n, 0)
}
