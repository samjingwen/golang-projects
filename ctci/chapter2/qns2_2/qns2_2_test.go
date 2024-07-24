package qns2_2

import "testing"
import . "ctci/chapter2"

type args struct {
	head *LinkedListNode
	k    int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"", args{nil, 0}, []int{}},
	{"", args{MakeNodeFromSlice(1), 1}, []int{1}},
	{"", args{MakeNodeFromSlice(1, 2), 1}, []int{2}},
	{"", args{MakeNodeFromSlice(1, 2), 2}, []int{1, 2}},
	{"", args{MakeNodeFromSlice(1, 2, 3, 4, 5), 2}, []int{4, 5}},
	{"", args{MakeNodeFromSlice(1, 2, 3, 4, 5), 1}, []int{5}},
}

func TestKthToLast1(t *testing.T) {
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			if got := KthToLast1(test.args.head, test.args.k); !got.Equals(test.want) {
				t.Errorf("KthToLast1() = %v, want %v", MakeSliceFromNode(got), test.want)
			}
		})
	}
}

func TestKthToLast2(t *testing.T) {
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			if got := KthToLast2(test.args.head, test.args.k); !got.Equals(test.want) {
				t.Errorf("KthToLast2() = %v, want %v", MakeSliceFromNode(got), test.want)
			}
		})
	}
}
