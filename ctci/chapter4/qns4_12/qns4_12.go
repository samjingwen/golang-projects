package qns4_12

import . "ctci/chapter4"

type void struct{}

var null = void{}

func PathsWithSum(node *TreeNode, targetSum int) int {
	var count int

	var iter func(node *TreeNode, sumSoFar int, sums map[int]void)
	iter = func(node *TreeNode, sumSoFar int, sums map[int]void) {
		if node == nil {
			return
		}
		sumSoFar += node.Data
		_, exists := sums[sumSoFar]
		if exists {
			count++
		}

		sums[sumSoFar+targetSum] = null

		iter(node.Left, sumSoFar, sums)
		iter(node.Right, sumSoFar, sums)

		delete(sums, sumSoFar+targetSum)
	}

	iter(node, 0, map[int]void{targetSum: null})
	return count
}
