package qns4_10

import "testing"
import . "ctci/chapter4"

func TestCheckSubtree(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{
			t1: &TreeNode{Data: 3,
				Left: &TreeNode{Data: 4}},
			t2: &TreeNode{Data: 3,
				Right: &TreeNode{Data: 4}}}, want: false},
		{args: args{
			t1: &TreeNode{Data: 3,
				Left: &TreeNode{Data: 3,
					Left: &TreeNode{Data: 4}},
				Right: &TreeNode{Data: 5}},
			t2: &TreeNode{Data: 3,
				Left:  &TreeNode{Data: 3},
				Right: &TreeNode{Data: 4}}}, want: false},
		{args: args{
			t1: &TreeNode{Data: 5,
				Left: &TreeNode{Data: 2,
					Left:  &TreeNode{Data: 1},
					Right: &TreeNode{Data: 3}},
				Right: &TreeNode{Data: 7,
					Left:  &TreeNode{Data: 6},
					Right: &TreeNode{Data: 8}}},
			t2: &TreeNode{Data: 7,
				Left:  &TreeNode{Data: 6},
				Right: &TreeNode{Data: 8}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSubtree(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("CheckSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSubArr(t *testing.T) {
	type args struct {
		arr    []int
		subArr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{2, 3, 0, 5, 1, 1, 2}, []int{3, 0, 5, 1}}, true},
		{"", args{[]int{2, 3, 3, 0, 5, 1, 1, 2}, []int{3, 0, 5, 1}}, true},
		{"", args{[]int{2, 3, 0, 5, 1, 1, 2}, []int{3, 0, 5, 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubArr(tt.args.arr, tt.args.subArr); got != tt.want {
				t.Errorf("findSubArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
