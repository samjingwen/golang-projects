package qns4_11

import (
	"math/rand"
	"time"
)

type RandomTree struct {
	rootNode *TreeNode
}

type TreeNode struct {
	data, size  int
	left, right *TreeNode
}

func (tree *RandomTree) insert(val int) {
	var iter func(node *TreeNode) *TreeNode
	iter = func(node *TreeNode) *TreeNode {
		if node == nil {
			return &TreeNode{data: val, size: 1}
		}
		if val == node.data {
			return node
		} else if val > node.data {
			node.right = iter(node.right)
			node.size++
		} else {
			node.left = iter(node.left)
			node.size++
		}
		return node
	}

	tree.rootNode = iter(tree.rootNode)
}

func (tree *RandomTree) find(val int) *TreeNode {
	var iter func(node *TreeNode) *TreeNode
	iter = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.data == val {
			return node
		} else if node.data > val {
			return iter(node.right)
		} else {
			return iter(node.left)
		}
	}
	return iter(tree.rootNode)
}

func (tree *RandomTree) delete(val int) {
	tree.rootNode = tree.rootNode.delete(val)
}

func (node *TreeNode) delete(val int) *TreeNode {
	var iter func(node *TreeNode) *TreeNode
	iter = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if val > node.data {
			node.right = iter(node.right)
		} else if val < node.data {
			node.left = iter(node.left)
		} else {
			if node.left == nil {
				return node.right
			} else if node.right == nil {
				return node.left
			} else {
				ptr := node.right
				for ptr.left != nil {
					ptr = ptr.left
				}
				node.data = ptr.data
				node.right = node.right.delete(ptr.data)
			}
		}
		return node
	}

	return iter(node)
}

func (tree *RandomTree) getRandomNode() *TreeNode {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	selected := random.Intn(tree.rootNode.size)

	var iter func(node *TreeNode, selected int) *TreeNode
	iter = func(node *TreeNode, selected int) *TreeNode {
		if node == nil {
			return nil
		}

		leftNodeSize := 0
		if node.left != nil {
			leftNodeSize = node.left.size
		}

		if selected < leftNodeSize {
			return iter(node.left, selected)
		} else if selected > leftNodeSize {
			return iter(node.right, selected-(leftNodeSize+1))
		} else {
			return node
		}
	}

	return iter(tree.rootNode, selected)
}

func (node *TreeNode) inOrderTraversal() []int {
	var res []int

	var iter func(node *TreeNode)
	iter = func(node *TreeNode) {
		if node == nil {
			return
		}
		iter(node.left)
		res = append(res, node.data)
		iter(node.right)
	}

	iter(node)
	return res
}
