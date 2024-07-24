package qns4_11

import (
	"fmt"
	"testing"
)

func TestRandomTree_insert(t *testing.T) {
	tree := &RandomTree{
		rootNode: &TreeNode{data: 5},
	}

	tree.insert(2)
	tree.insert(1)
	tree.insert(4)
	tree.insert(8)
	tree.insert(7)

	//        5
	//       / \
	//      /   \
	//     /     \
	//    /       \
	//   2         8
	//  / \       /
	// 1   4     7

	fmt.Printf("%+v\n", tree.rootNode)
	fmt.Printf("%+v\n", tree.rootNode.left)
	fmt.Printf("%+v\n", tree.rootNode.left.left)
	fmt.Printf("%+v\n", tree.rootNode.left.right)
	fmt.Printf("%+v\n", tree.rootNode.right)

	count := make(map[int]int)
	for i := 0; i < 10000; i++ {
		node := tree.getRandomNode()
		count[node.data]++
	}
	fmt.Printf("%v", count)
}

func TestTreeNode_delete(t *testing.T) {
	//           8
	//          / \
	//         /   \
	//        /     \
	//       /       \
	//      /         \
	//     /           \
	//    3            12
	//   / \           /
	//  /   \         /
	// 2     6       11
	//      / \
	//     5   7
	//    /
	//   4

	tree := &RandomTree{
		rootNode: &TreeNode{data: 8},
	}

	tree.insert(3)
	tree.insert(2)
	tree.insert(6)
	tree.insert(7)
	tree.insert(5)
	tree.insert(4)
	tree.insert(12)
	tree.insert(11)

	tree.delete(8)

	fmt.Println(tree.rootNode.inOrderTraversal())
}
