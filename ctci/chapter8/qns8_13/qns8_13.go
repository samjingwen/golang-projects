package qns8_13

import "sort"

// StackOfBoxes
// cuboids[i] = [width_i, length_i, height_i]
func StackOfBoxes(cuboids [][]int) int {
	sort.Sort(Cubes(cuboids))

	res := 0

	var iter func(index int, stack []int)
	iter = func(index int, stack []int) {
		if index == len(cuboids) {
			res = max(res, sum(stack, cuboids))
			return
		}

		for i := index; i < len(cuboids); i++ {
			curr := cuboids[i]
			if len(stack) > 0 {
				prev := cuboids[stack[len(stack)-1]]
				if curr[0] > prev[0] || curr[1] > prev[1] || curr[2] > prev[2] {
					continue
				}
			}

			stack = append(stack, i)
			iter(i+1, stack)
			stack = stack[:len(stack)-1]
		}
	}

	var stack []int
	iter(0, stack)
	return res
}

type Cubes [][]int

func (cubes Cubes) Less(i, j int) bool {
	if cubes[i][0] > cubes[j][0] {
		return true
	}
	return false
}

func (cubes Cubes) Swap(i, j int) {
	cubes[i], cubes[j] = cubes[j], cubes[i]
}

func (cubes Cubes) Len() int {
	return len(cubes)
}

func sum(stack []int, cuboids [][]int) int {
	res := 0
	for _, val := range stack {
		res += cuboids[val][2]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
