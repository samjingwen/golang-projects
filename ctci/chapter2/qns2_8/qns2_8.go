package qns2_8

import . "ctci/chapter2"

func LoopDetection(node *LinkedListNode) *LinkedListNode {
	fast, slow := node, node
	for {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	head := node
	for {
		if head == slow {
			return head
		}
		head = head.Next
		slow = slow.Next
	}
}
