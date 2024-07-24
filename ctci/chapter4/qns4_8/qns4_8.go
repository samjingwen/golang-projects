package qns4_8

import . "ctci/chapter4"

func FirstCommonAncestor(head, node1, node2 *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node == node1 || node == node2 {
			return node
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if left == node1 && right == node2 {
			return node
		}
		if left != nil {
			return left
		}
		if right != nil {
			return right
		}
		return nil
	}
	return dfs(head)
}
