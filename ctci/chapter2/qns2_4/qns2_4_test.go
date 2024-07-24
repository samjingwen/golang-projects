package qns2_4

import (
	"testing"
)
import . "ctci/chapter2"

type args struct {
	head      *LinkedListNode
	partition int
}

var tests = []struct {
	args args
	want [][]int
}{
	{args{MakeNodeFromSlice(3, 5, 8, 5, 10, 2, 1), 5},
		[][]int{{3, 2, 1, 5, 10, 5, 8}, {3, 2, 1, 5, 8, 5, 10}, {1, 2, 3, 5, 10, 5, 8}}},
	{args{MakeNodeFromSlice(5, 5, 5, 10, 1, 2, 3, 4), 5},
		[][]int{{1, 2, 3, 4, 5, 5, 5, 10}, {4, 3, 2, 1, 5, 5, 5, 10}}},
	{args{MakeNodeFromSlice(5, 5, 5, 10, 1, 2, 3, 4, 5), 5},
		[][]int{{1, 2, 3, 4, 5, 5, 5, 10, 5}, {4, 3, 2, 1, 5, 5, 5, 10, 5}}},
}

func TestPartition1(t *testing.T) {
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := Partition1(test.args.head, test.args.partition)
			success := false
			for _, expected := range test.want {
				if got.Equals(expected) {
					success = true
				}
			}
			if !success {
				t.Errorf("Partition1() = %v, want %v", MakeSliceFromNode(got), test.want)
			}
		})
	}
}

func TestPartition2(t *testing.T) {
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := Partition2(test.args.head, test.args.partition)
			success := false
			for _, expected := range test.want {
				if got.Equals(expected) {
					success = true
				}
			}
			if !success {
				t.Errorf("Partition2() = %v, want %v", MakeSliceFromNode(got), test.want)
			}
		})
	}
}

func TestPartition3(t *testing.T) {
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			got := Partition3(test.args.head, test.args.partition)
			success := false
			for _, expected := range test.want {
				if got.Equals(expected) {
					success = true
				}
			}
			if !success {
				t.Errorf("Partition3() = %v, want %v", MakeSliceFromNode(got), test.want)
			}
		})
	}
}
