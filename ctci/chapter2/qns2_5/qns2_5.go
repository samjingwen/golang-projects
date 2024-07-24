package qns2_5

import (
	. "ctci/chapter2"
)

// 1's is at head
func SumLists1(augend, addend *LinkedListNode) *LinkedListNode {
	result := &LinkedListNode{}
	runner := result

	i, j := augend, addend

	carry := 0
	for i != nil || j != nil {
		sum := &LinkedListNode{Data: carry}
		if i != nil {
			sum.Data += i.Data
			i = i.Next
		}
		if j != nil {
			sum.Data += j.Data
			j = j.Next
		}
		carry = sum.Data / 10
		sum.Data = sum.Data % 10

		runner.Next = sum
		runner = runner.Next
	}
	if carry != 0 {
		runner.Next = &LinkedListNode{Data: carry}
	}

	return result.Next
}

// 1's is at tail
func SumLists2(augend, addend *LinkedListNode) *LinkedListNode {
	var stack1 []int
	var stack2 []int
	for augend != nil || addend != nil {
		if augend != nil {
			stack1 = append(stack1, augend.Data)
			augend = augend.Next
		}
		if addend != nil {
			stack2 = append(stack2, addend.Data)
			addend = addend.Next
		}
	}

	var carry int
	var result *LinkedListNode
	for len(stack1) > 0 || len(stack2) > 0 {
		node := &LinkedListNode{Data: carry}
		if len(stack1) > 0 {
			node.Data += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			node.Data += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		carry = node.Data / 10
		node.Data %= 10

		node.Next = result
		result = node
	}
	if carry > 0 {
		node := &LinkedListNode{Data: carry}
		node.Next = result
		result = node
	}
	return result
}
