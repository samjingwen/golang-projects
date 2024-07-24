package qns4_3

import (
	. "ctci/chapter2"
	. "ctci/chapter3"
	. "ctci/chapter4"
)

// use level order traversal
// enqueue root node in queue
// deque all from queue, create linked lists, append to result
// enqueue child nodes into another temp queue
// set queue to temp
// repeat
func ListOfDepths(tree BST) []LinkedListNode {
	if tree.RootNode == nil {
		return nil
	}
	var res []LinkedListNode

	head := LinkedListNode{Data: tree.RootNode.Data}
	res = append(res, head)

	queue := Queue{}
	queue.Enqueue(tree.RootNode)
	for !queue.IsEmpty() {
		temp := Queue{}
		head = LinkedListNode{}
		ptr := &head
		for !queue.IsEmpty() {
			val, _ := queue.Deque()
			node := val.(*TreeNode)
			if node.Left != nil {
				ptr.Next = &LinkedListNode{Data: node.Left.Data}
				ptr = ptr.Next
				temp.Enqueue(node.Left)
			}
			if node.Right != nil {
				ptr.Next = &LinkedListNode{Data: node.Right.Data}
				ptr = ptr.Next
				temp.Enqueue(node.Right)
			}
		}
		if head.Next != nil {
			res = append(res, *head.Next)
		}
		queue = temp
	}
	return res
}

// depth first search
func ListOfDepthsRecursive(tree BST) []LinkedListNode {
	if tree.RootNode == nil {
		return nil
	}

	var iter func(*TreeNode, *[]LinkedListNode, int)
	iter = func(node *TreeNode, lists *[]LinkedListNode, level int) {
		if node == nil {
			return
		}
		if len(*lists) == level {
			list := LinkedListNode{Data: node.Data}
			*lists = append(*lists, list)
		} else {
			list := (*lists)[level]
			list.Append(node.Data)
			(*lists)[level] = list
		}
		iter(node.Left, lists, level+1)
		iter(node.Right, lists, level+1)
	}

	res := make([]LinkedListNode, 0)
	resPtr := &res
	iter(tree.RootNode, resPtr, 0)

	return res
}
