package qns3_1

import (
	"fmt"
)

type ThreeInOne []int

func (stack *ThreeInOne) Append(id int, val int) {
	var newStack ThreeInOne
	(*stack)[id]++
	lastIdx, _ := stack.getLastIdx(id)
	if len(*stack) < lastIdx+1 {
		newStack = make([]int, lastIdx+1)
		copy(newStack, *stack)
	}

	*stack = newStack
	(*stack)[lastIdx] = val
}

func (stack *ThreeInOne) Pop(id int) (int, error) {
	length := stack.getLength(id)
	if length == 0 {
		return 0, fmt.Errorf("stack with id: %v is empty", id)
	}
	lastIdx, _ := stack.getLastIdx(id)
	(*stack)[id]--
	return (*stack)[lastIdx], nil
}

func (stack *ThreeInOne) getLength(id int) int {
	return (*stack)[id]
}

func (stack *ThreeInOne) getLastIdx(id int) (int, error) {
	length := stack.getLength(id)
	if length == 0 {
		return 0, fmt.Errorf("stack with id: %v is empty", id)
	}
	return length*3 + id, nil
}
