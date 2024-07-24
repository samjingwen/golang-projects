package qns2_6

import . "ctci/chapter2"

func Palindrome(head *LinkedListNode) bool {
	ptr := head

	var iter func(*LinkedListNode) bool
	iter = func(node *LinkedListNode) bool {
		if node == nil {
			return true
		}

		res := iter(node.Next) && ptr.Data == node.Data
		ptr = ptr.Next
		return res
	}
	return iter(ptr)
}
