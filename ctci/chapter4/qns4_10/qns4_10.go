package qns4_10

import (
	. "ctci/chapter4"
	"math"
)

func CheckSubtree(t1, t2 *TreeNode) bool {
	arr1 := preOrderTraversal(t1)
	arr2 := preOrderTraversal(t2)
	return findSubArr(arr1, arr2)
}

// use max int as null symbol
func preOrderTraversal(node *TreeNode) []int {
	var res []int

	var iter func(node *TreeNode)
	iter = func(node *TreeNode) {
		if node == nil {
			res = append(res, math.MaxInt)
			return
		}
		res = append(res, node.Data)
		iter(node.Left)
		iter(node.Right)
	}

	iter(node)
	return res
}

func findSubArr(arr, subArr []int) bool {
	var i, j int

	for i < len(arr) && j < len(subArr) {
		if arr[i] == subArr[j] {
			i++
			j++
			if j >= len(subArr) {
				return true
			}
		} else {
			i = i - j + 1
			j = 0
		}
	}
	return false
}
