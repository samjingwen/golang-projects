package qns2_2

import . "ctci/chapter2"

func KthToLast1(head *LinkedListNode, k int) *LinkedListNode {
	length := 0
	curr := head
	for curr != nil {
		length += 1
		curr = curr.Next
	}

	times := length - k
	curr = head
	for i := 0; i < times; i++ {
		curr = curr.Next
	}
	return curr
}

// Recursive solution
func KthToLast2(head *LinkedListNode, k int) *LinkedListNode {
	var index int
	var found *LinkedListNode

	var iter func(*LinkedListNode)
	iter = func(curr *LinkedListNode) {
		if curr == nil {
			return
		}
		iter(curr.Next)
		index++
		if index == k {
			found = curr
		}
	}

	iter(head)
	return found
}
