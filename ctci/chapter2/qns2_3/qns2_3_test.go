package qns2_3

import (
	. "ctci/chapter2"
	"testing"
)

type args struct {
	head     *LinkedListNode
	toDelete *LinkedListNode
}

var tests = []struct {
	args args
	want []int
}{
	{makeArgs([]int{1, 2, 3}, 1), []int{1, 3}},
	{makeArgs([]int{1, 2, 3, 4}, 2), []int{1, 2, 4}},
	{makeArgs([]int{1, 2, 3, 4, 5, 6, 7}, 2), []int{1, 2, 4, 5, 6, 7}},
}

func TestDeleteMiddleNode(t *testing.T) {
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			DeleteMiddleNode(test.args.toDelete)
			if !test.args.head.Equals(test.want) {
				t.Errorf("DeleteMiddleNode() = %v, want %v", test.args, test.want)
			}
		})
	}
}

func makeArgs(arr []int, index int) args {
	head := MakeNodeFromSlice(arr...)

	ptr := head
	for i := 0; i < index; i++ {
		ptr = ptr.Next
	}
	return args{head, ptr}
}
