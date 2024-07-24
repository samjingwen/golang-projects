package qns2_1

import . "ctci/chapter2"

func RemoveDups1(head *LinkedListNode) {
	if head == nil {
		return
	}
	cache := make(map[int]bool)
	cache[head.Data] = true
	ptr := head
	for ptr != nil && ptr.Next != nil {
		_, exists := cache[ptr.Next.Data]
		if exists {
			ptr.Next = ptr.Next.Next
			continue
		} else {
			cache[ptr.Next.Data] = true
		}
		ptr = ptr.Next
	}
}

func RemoveDups2(head *LinkedListNode) {
	curr := head
	for curr != nil {
		prev := curr
		ptr := curr.Next
		for ptr != nil {
			if ptr.Data == curr.Data {
				prev.Next = prev.Next.Next
			}
			prev = ptr
			ptr = ptr.Next
		}
		curr = curr.Next
	}
}
