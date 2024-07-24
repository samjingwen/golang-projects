package qns4_2

import (
	. "ctci/chapter4"
)

func MinimalTree(arr []int) BST {
	var iter func(*TreeNode, []int) *TreeNode
	iter = func(node *TreeNode, arr []int) *TreeNode {
		if len(arr) > 0 {
			mid := len(arr) / 2
			node = &TreeNode{Data: arr[mid]}
			node.Left = iter(node.Left, arr[:mid])
			node.Right = iter(node.Right, arr[mid+1:])
			return node
		}
		return nil
	}
	bst := BST{}
	bst.RootNode = iter(bst.RootNode, arr)
	return bst
}
