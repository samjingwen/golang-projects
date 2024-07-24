package chapter4

import (
	"reflect"
	"testing"
)

func TestBST(t *testing.T) {
	bst := BST{}

	//     5
	//    / \
	//   3   7
	//  / \   \
	// 2   4   8
	bst.Add(5)
	bst.Add(3)
	bst.Add(4)
	bst.Add(2)
	bst.Add(7)
	bst.Add(8)

	depth := bst.Depth()
	if depth != 3 {
		t.Errorf("depth = %v, expected = %v", depth, 3)
	}

	arr := bst.LevelOrderTraversal()
	if expected := []int{5, 3, 7, 2, 4, 8}; !reflect.DeepEqual(arr, expected) {
		t.Errorf("levelOrderTraversal = %v, expected = %v", arr, expected)
	}

	arr = bst.PreOrderTraversal()
	if expected := []int{5, 3, 2, 4, 7, 8}; !reflect.DeepEqual(arr, expected) {
		t.Errorf("levelOrderTraversal = %v, expected = %v", arr, expected)
	}

	arr = bst.PostOrderTraversal()
	if expected := []int{2, 4, 3, 8, 7, 5}; !reflect.DeepEqual(arr, expected) {
		t.Errorf("levelOrderTraversal = %v, expected = %v", arr, expected)
	}
}
