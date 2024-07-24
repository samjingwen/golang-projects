package qns4_2

import (
	"testing"
)

func TestMinimalTree(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{1}}, 1},
		{"", args{[]int{1, 2}}, 2},
		{"", args{[]int{1, 2, 3}}, 2},
		{"", args{[]int{1, 2, 3, 4}}, 3},
		{"", args{[]int{1, 2, 3, 4, 5, 6, 7}}, 3},
		{"", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}}, 4},
		{"", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}}, 4},
		{"", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimalTree(tt.args.arr); got.Depth() != tt.want {
				t.Errorf("MinimalTree() = %v, depth %v, want %v", got.LevelOrderTraversal(), got.Depth(), tt.want)
			}
		})
	}
}
