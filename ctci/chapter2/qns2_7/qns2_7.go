package qns2_7

import (
	. "ctci/chapter2"
)

func Intersection(node1, node2 *LinkedListNode) *LinkedListNode {
	len1 := getLen(node1)
	len2 := getLen(node2)
	diff := abs(len1, len2)

	var helper func(*LinkedListNode, *LinkedListNode, int, int) *LinkedListNode
	helper = func(slow *LinkedListNode, fast *LinkedListNode, x, y int) *LinkedListNode {
		for i := 0; i < diff; i++ {
			fast = fast.Next
		}

		for {
			if slow == fast {
				return slow
			}
			slow = slow.Next
			fast = fast.Next
		}
	}

	if len1 > len2 {
		return helper(node2, node1, len2, len1)
	} else {
		return helper(node1, node2, len1, len2)
	}
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func getLen(node *LinkedListNode) int {
	ptr := node
	length := 0
	for ptr != nil {
		length++
		ptr = ptr.Next
	}
	return length
}
