package chapter4

import (
	. "ctci/chapter3"
)

type BST struct {
	Size     int
	RootNode *TreeNode
}

type TreeNode struct {
	Data        int
	Left, Right *TreeNode
}

func (tree *BST) Add(elem int) {
	var iter func(node *TreeNode, elem int) *TreeNode
	iter = func(node *TreeNode, elem int) *TreeNode {
		if node == nil {
			return &TreeNode{Data: elem}
		}
		if elem == node.Data {
			return node
		} else if elem < node.Data {
			node.Left = iter(node.Left, elem)
		} else {
			node.Right = iter(node.Right, elem)
		}
		return node
	}
	tree.RootNode = iter(tree.RootNode, elem)
	tree.Size++
}

func (tree *BST) Contains(elem int) bool {
	var iter func(*TreeNode, int) bool
	iter = func(node *TreeNode, elem int) bool {
		if node == nil {
			return false
		}
		if elem == node.Data {
			return true
		} else if elem < node.Data {
			return iter(node.Left, elem)
		} else {
			return iter(node.Right, elem)
		}
	}
	return iter(tree.RootNode, elem)
}

func (tree *BST) Depth() int {
	var iter func(*TreeNode) int
	iter = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(iter(node.Left), iter(node.Right)) + 1
	}
	return iter(tree.RootNode)
}

func (tree *BST) LevelOrderTraversal() []int {
	return tree.RootNode.LevelOrderTraversal()
}

func (node *TreeNode) LevelOrderTraversal() []int {
	var res []int

	queue := Queue{}
	queue.Enqueue(node)
	for !queue.IsEmpty() {
		curr, _ := queue.Deque()
		node := curr.(*TreeNode)
		res = append(res, node.Data)
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}
	return res
}

func (tree *BST) PreOrderTraversal() []int {
	return tree.RootNode.PreOrderTraversal()
}

func (node *TreeNode) PreOrderTraversal() []int {
	var res []int

	var iter func(node *TreeNode)
	iter = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Data)
		iter(node.Left)
		iter(node.Right)
	}

	iter(node)
	return res
}

func (tree *BST) PostOrderTraversal() []int {
	return tree.RootNode.PostOrderTraversal()
}

func (node *TreeNode) PostOrderTraversal() []int {
	var res []int

	var iter func(node *TreeNode)
	iter = func(node *TreeNode) {
		if node == nil {
			return
		}
		iter(node.Left)
		iter(node.Right)
		res = append(res, node.Data)
	}

	iter(node)
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
