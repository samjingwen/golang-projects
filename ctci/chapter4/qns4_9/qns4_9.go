package qns4_9

import . "ctci/chapter4"

func BSTSequences(node *TreeNode) [][]int {
	return helper(node, [][]int{})
}

func helper(node *TreeNode, res [][]int) [][]int {
	if node == nil {
		return [][]int{}
	}

	left := helper(node.Left, res)
	right := helper(node.Right, res)

	if len(left) == 0 && len(right) == 0 {
		res = merge([]int{}, []int{}, []int{node.Data}, res)
	} else if len(left) == 0 {
		for _, eachRight := range right {
			res = merge([]int{}, eachRight, []int{node.Data}, res)
		}
	} else if len(right) == 0 {
		for _, eachLeft := range left {
			res = merge([]int{}, eachLeft, []int{node.Data}, res)
		}
	} else if len(left) > 0 && len(right) > 0 {
		for _, eachLeft := range left {
			for _, eachRight := range right {
				res = merge(eachLeft, eachRight, []int{node.Data}, res)
			}
		}
	}
	return res
}

func merge(left, right, prefix []int, res [][]int) [][]int {
	if len(left) == 0 && len(right) == 0 && len(prefix) > 0 {
		temp := make([]int, len(prefix))
		copy(temp, prefix)
		res = append(res, temp)
	}
	if len(left) > 0 {
		res = merge(left[1:], right, append(prefix, left[0]), res)
	}
	if len(right) > 0 {
		res = merge(left, right[1:], append(prefix, right[0]), res)
	}
	return res
}
