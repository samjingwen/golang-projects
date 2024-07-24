package qns3_1

import (
	"reflect"
	"testing"
)

func TestThreeInOne(t *testing.T) {
	stack := ThreeInOne{0, 0, 0}
	stack.Append(0, 1)

	expected := []int{0, 0, 0, 1}
	if reflect.DeepEqual(stack, expected) {
		t.Errorf("actual = %v, expected %v", stack, expected)
	}

	stack.Append(1, 2)
	stack.Append(2, 3)
	expected = []int{0, 0, 0, 1, 2, 0, 0, 3, 0}
	if reflect.DeepEqual(stack, expected) {
		t.Errorf("actual = %v, expected %v", stack, expected)
	}

	val, _ := stack.Pop(0)
	expected = []int{0, 0, 0, 0, 2, 0, 0, 3, 0}
	if val != 1 && reflect.DeepEqual(stack, expected) {
		t.Errorf("actual = %v, expected %v", stack, expected)
	}

	val, _ = stack.Pop(1)
	expected = []int{0, 0, 0, 0, 2, 0, 0, 0, 0}
	if val != 3 && reflect.DeepEqual(stack, expected) {
		t.Errorf("actual = %v, expected %v", stack, expected)
	}
}
