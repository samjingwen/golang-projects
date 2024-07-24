package common

import (
	"testing"
)

func TestQueue(t *testing.T) {
	var queue Queue[int]
	queue.Enqueue(1)
	queue.Enqueue(5)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if res := queue.IsEmpty(); res {
		t.Errorf("should be not empty")
	}
	if res := queue.Size(); res != 4 {
		t.Errorf("should be 4")
	}
	if res, _ := queue.Deque(); res != 1 {
		t.Errorf("should be 1")
	}
	if res, _ := queue.Deque(); res != 5 {
		t.Errorf("should be 5")
	}
	if res, _ := queue.Deque(); res != 2 {
		t.Errorf("should be 2")
	}
	if res, _ := queue.Deque(); res != 3 {
		t.Errorf("should be 3")
	}
	if _, err := queue.Deque(); err.Error() != "queue is empty" {
		t.Errorf("should be \"queue is empty\"")
	}
	if res := queue.IsEmpty(); !res {
		t.Errorf("should be empty")
	}

}
