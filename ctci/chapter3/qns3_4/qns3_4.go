package qns3_4

import (
	. "ctci/chapter3"
)

type MyQueue struct {
	front, back Stack
}

func (queue *MyQueue) Enqueue(value int) {
	for !queue.back.IsEmpty() {
		pop, _ := queue.back.Pop()
		queue.front.Push(pop)
	}
	queue.front.Push(value)

	for !queue.front.IsEmpty() {
		pop, _ := queue.front.Pop()
		queue.back.Push(pop)
	}
}

func (queue *MyQueue) Deque() int {
	res, _ := queue.back.Pop()
	return res.(int)
}
