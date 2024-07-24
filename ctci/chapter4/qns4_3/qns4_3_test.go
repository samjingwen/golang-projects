package qns4_3

import (
	. "ctci/chapter2"
	. "ctci/chapter4"
	"reflect"
	"testing"
)

type args struct {
	tree BST
}

var tests = []struct {
	name string
	args args
	want []LinkedListNode
}{
	{args: args{tree: BST{}}, want: nil},
	{args: args{tree: BST{
		Size:     1,
		RootNode: &TreeNode{Data: 1}}},
		want: []LinkedListNode{{Data: 1}}},
	{args: args{BST{
		Size: 3,
		RootNode: &TreeNode{
			Data:  1,
			Left:  &TreeNode{Data: 2},
			Right: &TreeNode{Data: 3}}}},
		want: []LinkedListNode{
			{Data: 1},
			{Data: 2, Next: &LinkedListNode{Data: 3}}}},
	{args: args{tree: BST{
		Size: 4,
		RootNode: &TreeNode{
			Data:  1,
			Left:  &TreeNode{Data: 2, Left: &TreeNode{Data: 5}},
			Right: &TreeNode{Data: 3}}}},
		want: []LinkedListNode{
			{Data: 1},
			{Data: 2, Next: &LinkedListNode{Data: 3}},
			{Data: 5}}},
}

func TestListOfDepths(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListOfDepths(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListOfDepths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListOfDepthsRecursive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListOfDepthsRecursive(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListOfDepthsRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
