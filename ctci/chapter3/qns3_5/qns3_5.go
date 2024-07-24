package qns3_5

import (
	. "ctci/chapter3"
)

// SortStack sort stack using only another stack
func SortStack(stack *Stack) {
	// find max by popping from orig stack and pushing to temp stack
	// pop and push non-sorted segments back to orig stack
	// push max to temp stack
	// repeat
	length := stack.Size()
	temp := Stack{}
	for i := 0; i < length; i++ {
		maxPop, _ := stack.Pop()
		maxVal := maxPop.(int)
		for j := 0; j < length-i-1; j++ {
			pop, _ := stack.Pop()
			val := pop.(int)
			if val > maxVal {
				temp.Push(maxVal)
				maxVal = val
			} else {
				temp.Push(val)
			}
		}

		for j := 0; j < length-i-1; j++ {
			val, _ := temp.Pop()
			stack.Push(val)
		}
		temp.Push(maxVal)
	}
	*stack = temp
}
