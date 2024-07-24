package qns4_6

import (
	"testing"
)

func TestSuccessor(t *testing.T) {

	node5 := &TreeNode{Data: 5}
	node1 := &TreeNode{Data: 1}
	node0 := &TreeNode{Data: 0}
	node2 := &TreeNode{Data: 2}
	node3 := &TreeNode{Data: 3}
	node4 := &TreeNode{Data: 4}
	node9 := &TreeNode{Data: 9}
	node6 := &TreeNode{Data: 6}
	node7 := &TreeNode{Data: 7}
	node8 := &TreeNode{Data: 8}
	node10 := &TreeNode{Data: 10}

	node5.Left = node1
	node5.Right = node9

	node1.Parent = node5
	node1.Left = node0
	node1.Right = node2

	node0.Parent = node1

	node2.Parent = node1
	node2.Right = node3

	node3.Parent = node2
	node3.Right = node4

	node4.Parent = node3

	node9.Parent = node5
	node9.Left = node6
	node9.Right = node10

	node10.Parent = node9

	node6.Parent = node9
	node6.Right = node7

	node7.Parent = node6
	node7.Right = node8

	node8.Parent = node7

	if got := Successor(node2); got.Data != 3 {
		t.Errorf("Successor() = %v, want %v", got.Data, 3)
	}

	if got := Successor(node5); got.Data != 6 {
		t.Errorf("Successor() = %v, want %v", got.Data, 6)
	}

	if got := Successor(node8); got.Data != 9 {
		t.Errorf("Successor() = %v, want %v", got.Data, 9)
	}

	if got := Successor(node1); got.Data != 2 {
		t.Errorf("Successor() = %v, want %v", got.Data, 2)
	}

	if got := Successor(node4); got.Data != 5 {
		t.Errorf("Successor() = %v, want %v", got.Data, 5)
	}
}
