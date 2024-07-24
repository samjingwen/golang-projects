package qns4_5

import (
	. "ctci/chapter4"
	"math"
)

func ValidateBST(node *TreeNode) bool {
	if node == nil {
		return true
	}
	var iter func(node *TreeNode, min, max int) bool
	iter = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		return node.Data < max && node.Data > min &&
			iter(node.Left, min, node.Data) &&
			iter(node.Right, node.Data, max)
	}
	return iter(node.Left, math.MinInt, node.Data) &&
		iter(node.Right, node.Data, math.MaxInt)
}
