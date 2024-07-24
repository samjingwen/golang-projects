package qns3_2

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := Stack{}
	stack.append(1)
	stack.append(2)
	stack.append(3)

	expected := stack.min()
	if expected != 1 {
		t.Errorf("actual = %v, expected %v", 1, expected)
	}

	stack.append(0)
	expected = stack.min()
	if expected != 0 {
		t.Errorf("actual = %v, expected %v", 0, expected)
	}

	expected = stack.pop()
	if expected != 0 {
		t.Errorf("actual = %v, expected %v", 0, expected)
	}

	expected = stack.min()
	if expected != 1 {
		t.Errorf("actual = %v, expected %v", 1, expected)
	}
}
