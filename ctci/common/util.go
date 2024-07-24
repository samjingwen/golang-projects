package common

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func ElementsAreEqual[T comparable](xs, ys []T) bool {
	if len(xs) != len(ys) {
		return false
	}

	diff := make(map[T]int, len(xs))
	for _, x := range xs {
		diff[x]++
	}

	for _, y := range ys {
		diff[y]--
	}

	for _, count := range diff {
		if count < 0 {
			return false
		}
	}
	return true
}
