package qns4_4

import . "ctci/chapter4"

func CheckBalanced(tree *TreeNode) bool {
	if tree == nil {
		return true
	}

	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(depth(node.Left), depth(node.Right)) + 1
	}
	return abs(depth(tree.Left), depth(tree.Right)) < 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
