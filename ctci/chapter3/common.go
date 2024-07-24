package chapter3

import "errors"

type Stack []interface{}

func NewStack(arr ...interface{}) *Stack {
	return (*Stack)(&arr)
}

func (stack *Stack) Push(val interface{}) {
	*stack = append(*stack, val)
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	val := (*stack)[stack.Size()-1]
	*stack = (*stack)[:stack.Size()-1]
	return val, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return (*stack)[stack.Size()-1], nil
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Size() int {
	return len(*stack)
}

func (stack *Stack) Values() []interface{} {
	return *stack
}

type Queue []interface{}

func (queue *Queue) Enqueue(val interface{}) {
	*queue = append(*queue, val)
}

func (queue *Queue) Deque() (interface{}, error) {
	if len(*queue) == 0 {
		return 0, errors.New("queue is empty")
	}
	val := (*queue)[0]
	*queue = (*queue)[1:]
	return val, nil
}

func (queue *Queue) Peek() (interface{}, error) {
	if len(*queue) == 0 {
		return 0, errors.New("queue is empty")
	}
	return (*queue)[0], nil
}

func (queue *Queue) Size() int {
	return len(*queue)
}

func (queue *Queue) IsEmpty() bool {
	return len(*queue) == 0
}
