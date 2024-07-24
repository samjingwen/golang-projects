package qns4_5

import "testing"
import . "ctci/chapter4"

func TestValidateBST(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{nil}, want: true},
		{args: args{node: &TreeNode{}}, want: true},
		{args: args{
			node: &TreeNode{Data: 5,
				Left:  &TreeNode{Data: 1},
				Right: &TreeNode{Data: 2}}}, want: false},
		{args: args{
			node: &TreeNode{Data: 5,
				Left:  &TreeNode{Data: 1},
				Right: &TreeNode{Data: 9}}}, want: true},
		{args: args{
			node: &TreeNode{Data: 5,
				Left: &TreeNode{Data: 1,
					Right: &TreeNode{Data: 3}},
				Right: &TreeNode{Data: 9}}}, want: true},
		{args: args{
			node: &TreeNode{Data: 5,
				Left: &TreeNode{Data: 1,
					Right: &TreeNode{Data: 6}},
				Right: &TreeNode{Data: 9}}}, want: false},
		{args: args{
			node: &TreeNode{Data: 5,
				Left: &TreeNode{Data: 1,
					Right: &TreeNode{Data: 4}},
				Right: &TreeNode{Data: 9}}}, want: true},
		{args: args{
			node: &TreeNode{Data: 5,
				Left: &TreeNode{Data: 1},
				Right: &TreeNode{Data: 9,
					Left: &TreeNode{Data: 4}}}}, want: false},
		{args: args{
			node: &TreeNode{Data: 5,
				Left: &TreeNode{Data: 1},
				Right: &TreeNode{Data: 9,
					Left: &TreeNode{Data: 6}}}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateBST(tt.args.node); got != tt.want {
				t.Errorf("ValidateBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
