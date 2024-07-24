package qns8_3

import . "ctci/common"

func MagicIndex(arr []int) int {
	if len(arr) <= 0 {
		return -1
	}

	mid := len(arr) / 2
	if arr[mid] > mid {
		return MagicIndex(arr[:mid])
	} else if arr[mid] < mid {
		return MagicIndex(arr[mid+1:])
	} else {
		return mid
	}
}
func MagicIndexUnique(arr []int) int {
	var iter func(start, end int) int
	iter = func(start, end int) int {
		if end < start {
			return -1
		}

		mid := (start + end) / 2
		if mid == arr[mid] {
			return mid
		}

		right := iter(Max(arr[mid], mid+1), end)
		if right >= 0 {
			return right
		}

		left := iter(start, Min(arr[mid], mid-1))
		return left

	}

	return iter(0, len(arr)-1)
}
