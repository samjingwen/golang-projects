package qns8_4

func PowerSet(set []int) [][]int {
	var iter func(index int) [][]int
	iter = func(index int) [][]int {
		var subsets [][]int

		if index == len(set) {
			subsets = append(subsets, []int{})
		} else {
			allSubsets := iter(index + 1)
			subsets = append(subsets, allSubsets...)

			for _, currSet := range allSubsets {
				newSet := make([]int, len(currSet))
				copy(newSet, currSet)
				newSet = append(newSet, set[index])

				subsets = append(subsets, newSet)
			}
		}
		return subsets
	}

	return iter(0)
}
