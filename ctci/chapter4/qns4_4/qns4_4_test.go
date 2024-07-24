package qns4_4

import (
	. "ctci/chapter4"
	"testing"
)

func TestCheckBalanced(t *testing.T) {
	type args struct {
		tree *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{nil}, true},
		{"", args{&TreeNode{}}, true},
		{"", args{
			&TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{}},
				Right: &TreeNode{}}}, true},
		{"", args{
			&TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Right: &TreeNode{}}},
				Right: &TreeNode{}}}, false},
		{"", args{
			&TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{}}}}, false},
		{"", args{
			&TreeNode{
				Left: &TreeNode{
					Left:  &TreeNode{},
					Right: &TreeNode{}}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckBalanced(tt.args.tree); got != tt.want {
				t.Errorf("CheckBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
