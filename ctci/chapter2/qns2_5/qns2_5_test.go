package qns2_5

import (
	"reflect"
	"testing"
)
import . "ctci/chapter2"

func TestSumLists1(t *testing.T) {
	type args struct {
		augend *LinkedListNode
		addend *LinkedListNode
	}
	tests := []struct {
		name string
		args args
		want *LinkedListNode
	}{
		{"", args{MakeNodeFromSlice(7, 1, 6), MakeNodeFromSlice(5, 9, 2)}, MakeNodeFromSlice(2, 1, 9)},
		{"", args{MakeNodeFromSlice(7, 1, 6), MakeNodeFromSlice(5, 9, 3)}, MakeNodeFromSlice(2, 1, 0, 1)},
		{"", args{MakeNodeFromSlice(7, 1, 6), MakeNodeFromSlice(5, 9, 3, 9, 7)}, MakeNodeFromSlice(2, 1, 0, 0, 8)},
		{"", args{MakeNodeFromSlice(7, 1, 6, 2, 3), MakeNodeFromSlice(5, 9, 3)}, MakeNodeFromSlice(2, 1, 0, 3, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumLists1(tt.args.augend, tt.args.addend); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumLists1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumLists2(t *testing.T) {
	type args struct {
		augend *LinkedListNode
		addend *LinkedListNode
	}
	tests := []struct {
		name string
		args args
		want *LinkedListNode
	}{
		{"", args{MakeNodeFromSlice(6, 1, 7), MakeNodeFromSlice(2, 9, 5)}, MakeNodeFromSlice(9, 1, 2)},
		{"", args{MakeNodeFromSlice(6, 1, 7), MakeNodeFromSlice(3, 9, 5)}, MakeNodeFromSlice(1, 0, 1, 2)},
		{"", args{MakeNodeFromSlice(6, 1, 7), MakeNodeFromSlice(7, 9, 3, 9, 5)}, MakeNodeFromSlice(8, 0, 0, 1, 2)},
		{"", args{MakeNodeFromSlice(3, 2, 6, 1, 7), MakeNodeFromSlice(3, 9, 5)}, MakeNodeFromSlice(3, 3, 0, 1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumLists2(tt.args.augend, tt.args.addend); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumLists2() = %v, want %v", got, tt.want)
			}
		})
	}
}
