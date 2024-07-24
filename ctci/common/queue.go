package common

import "errors"

type Queue[T any] []T

func (queue *Queue[T]) Enqueue(val T) {
	*queue = append(*queue, val)
}

func (queue *Queue[T]) Deque() (T, error) {
	if len(*queue) == 0 {
		var t T
		return t, errors.New("queue is empty")
	}
	val := (*queue)[0]
	*queue = (*queue)[1:]
	return val, nil
}

func (queue *Queue[T]) Peek() (T, error) {
	if len(*queue) == 0 {
		var t T
		return t, errors.New("queue is empty")
	}
	return (*queue)[0], nil
}

func (queue *Queue[T]) Size() int {
	return len(*queue)
}

func (queue *Queue[T]) IsEmpty() bool {
	return len(*queue) == 0
}

func (queue *Queue[T]) IsNotEmpty() bool {
	return !queue.IsEmpty()
}
