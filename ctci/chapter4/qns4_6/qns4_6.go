package qns4_6

func Successor(node *TreeNode) *TreeNode {
	if node.Right == nil {
		if node.Parent == nil {
			return nil
		}
		parentNode := node.Parent
		for parentNode.Right == node {
			node = parentNode
			parentNode = parentNode.Parent
		}
		return parentNode
	}
	return getLeftMostLeafNode(node.Right)
}

func getLeftMostLeafNode(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node
	}
	return getLeftMostLeafNode(node.Left)
}

type TreeNode struct {
	Data                int
	Parent, Left, Right *TreeNode
}
