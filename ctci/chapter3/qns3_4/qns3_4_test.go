package qns3_4

import "testing"

func TestMyQueue(t *testing.T) {
	queue := MyQueue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	actual := queue.Deque()
	if actual != 1 {
		t.Errorf("actual = %v, expected = %v", actual, 1)
	}

	actual = queue.Deque()
	if actual != 2 {
		t.Errorf("actual = %v, expected = %v", actual, 2)
	}

	queue.Enqueue(4)
	actual = queue.Deque()
	if actual != 3 {
		t.Errorf("actual = %v, expected = %v", actual, 3)
	}

	actual = queue.Deque()
	if actual != 4 {
		t.Errorf("actual = %v, expected = %v", actual, 4)
	}
}
