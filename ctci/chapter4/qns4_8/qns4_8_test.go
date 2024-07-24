package qns4_8

import (
	"reflect"
	"testing"
)
import . "ctci/chapter4"

type args struct {
	head  *TreeNode
	node1 *TreeNode
	node2 *TreeNode
}

func TestFirstCommonAncestor(t *testing.T) {
	args1, want1 := makeArgs1()
	args2, want2 := makeArgs2()
	args3, want3 := makeArgs3()
	args4, want4 := makeArgs4()
	args5, want5 := makeArgs5()

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"", args1, want1},
		{"", args2, want2},
		{"", args3, want3},
		{"", args4, want4},
		{"", args5, want5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstCommonAncestor(tt.args.head, tt.args.node1, tt.args.node2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstCommonAncestor() = %v, want %v", got.Data, tt.want.Data)
			}
		})
	}
}

//		  1
//		 / \
//		/   \
//	   /     \
//	  2       3
//	 / \     / \
//
// 4   5   6   7
//
//	 /
//	8
func makeArgs1() (args, *TreeNode) {
	node1 := &TreeNode{Data: 1}
	node2 := &TreeNode{Data: 2}
	node3 := &TreeNode{Data: 3}
	node4 := &TreeNode{Data: 4}
	node5 := &TreeNode{Data: 5}
	node6 := &TreeNode{Data: 6}
	node7 := &TreeNode{Data: 7}
	node8 := &TreeNode{Data: 8}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node6.Left = node8

	return args{node1, node4, node8}, node1
}

func makeArgs2() (args, *TreeNode) {
	node1 := &TreeNode{Data: 1}
	node2 := &TreeNode{Data: 2}
	node3 := &TreeNode{Data: 3}
	node4 := &TreeNode{Data: 4}
	node5 := &TreeNode{Data: 5}
	node6 := &TreeNode{Data: 6}
	node7 := &TreeNode{Data: 7}
	node8 := &TreeNode{Data: 8}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node6.Left = node8

	return args{node1, node8, node7}, node3
}

func makeArgs3() (args, *TreeNode) {
	node1 := &TreeNode{Data: 1}
	node2 := &TreeNode{Data: 2}
	node3 := &TreeNode{Data: 3}
	node4 := &TreeNode{Data: 4}
	node5 := &TreeNode{Data: 5}
	node6 := &TreeNode{Data: 6}
	node7 := &TreeNode{Data: 7}
	node8 := &TreeNode{Data: 8}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node6.Left = node8

	return args{node1, node4, node5}, node2
}

func makeArgs4() (args, *TreeNode) {
	node1 := &TreeNode{Data: 1}
	node2 := &TreeNode{Data: 2}
	node3 := &TreeNode{Data: 3}
	node4 := &TreeNode{Data: 4}
	node5 := &TreeNode{Data: 5}
	node6 := &TreeNode{Data: 6}
	node7 := &TreeNode{Data: 7}
	node8 := &TreeNode{Data: 8}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node6.Left = node8

	return args{node1, node4, node2}, node2
}

func makeArgs5() (args, *TreeNode) {
	node1 := &TreeNode{Data: 1}
	node2 := &TreeNode{Data: 2}
	node3 := &TreeNode{Data: 3}
	node4 := &TreeNode{Data: 4}
	node5 := &TreeNode{Data: 5}
	node6 := &TreeNode{Data: 6}
	node7 := &TreeNode{Data: 7}
	node8 := &TreeNode{Data: 8}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node6.Left = node8

	return args{node1, node1, node1}, node1
}
