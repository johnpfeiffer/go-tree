package gotree

import (
	"fmt"
	"strconv"
	"testing"
)

// Manually building each tree variation once
var BinaryTreeTestCases = []struct {
	a                 []int
	tree              *BinaryTree
	preOrderTraversal string
	inOrderTraversal  string
	height            int
}{
	{a: nil, tree: &BinaryTree{}, preOrderTraversal: "", inOrderTraversal: "", height: 0},
	{a: []int{}, tree: &BinaryTree{}, preOrderTraversal: "", inOrderTraversal: "", height: 0},
	{a: []int{1}, tree: &BinaryTree{Root: &Node{Data: 1}},
		preOrderTraversal: "1 ", inOrderTraversal: "1", height: 0},
	{a: []int{1, 2}, tree: &BinaryTree{Root: &Node{Data: 1, Left: &Node{Data: 2}}},
		preOrderTraversal: "1 2 ", inOrderTraversal: "2 1", height: 1},
	// perfect tree /\
	{a: []int{1, 2, 3}, tree: &BinaryTree{Root: &Node{Data: 1, Left: &Node{Data: 2}, Right: &Node{Data: 3}}},
		preOrderTraversal: "1 2 3 ", inOrderTraversal: "2 1 3", height: 1},
	// TODO
	// degenerate trees  / / \ \
	//                  /  \ /  \
	// {dataValues: []int{1, 2, -1}, tree: BinarySearchTree{Root: &Node{Data: 2, Left: &Node{Data: 1, Left: &Node{Data: -1}}}},
	// 	preOrderTraversal: "2 1 -1 ", height: 2},
	// {dataValues: []int{2, 0, 1}, tree: BinarySearchTree{Root: &Node{Data: 2, Left: &Node{Data: 0, Right: &Node{Data: 1}}}},
	// 	preOrderTraversal: "2 0 1 ", height: 2},
	// {dataValues: []int{2, 5, 4}, tree: BinarySearchTree{Root: &Node{Data: 2, Right: &Node{Data: 5, Left: &Node{Data: 4}}}},
	// 	preOrderTraversal: "2 5 4 ", height: 2},
	// {dataValues: []int{2, 5, 6}, tree: BinarySearchTree{Root: &Node{Data: 2, Right: &Node{Data: 5, Right: &Node{Data: 6}}}},
	// 	preOrderTraversal: "2 5 6 ", height: 2},

}

func TestCreateBinarySubTreeStringsSuccess(t *testing.T) {
	var testCases = []struct {
		a              []string
		index          int
		expectedHeight int
	}{
		{a: []string{"1"}, index: 0, expectedHeight: 1},
		{a: []string{"1", "nil"}, index: 0, expectedHeight: 1},
		{a: []string{"1", "2"}, index: 0, expectedHeight: 2},
		{a: []string{"1", "2", "nil", "nil"}, index: 0, expectedHeight: 2},
		{a: []string{"1", "2", "3"}, index: 0, expectedHeight: 2},
		{a: []string{"1", "nil", "3"}, index: 0, expectedHeight: 2},
		{a: []string{"1", "2", "nil", "3"}, index: 0, expectedHeight: 3},
		{a: []string{"1", "2", "nil", "3", "4"}, index: 0, expectedHeight: 3},
		{a: []string{"1", "nil", "3", "nil", "nil", "6"}, index: 0, expectedHeight: 3},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.a), func(t *testing.T) {
			root := CreateBinarySubtreeStrings(tc.a, tc.index)
			expectedRootData, _ := strconv.Atoi(tc.a[0])
			if expectedRootData != root.Data {
				t.Errorf("Expected root value: %d but received %d", expectedRootData, root.Data)
			}
			result := SubtreeHeight(root)
			if tc.expectedHeight != result {
				t.Errorf("Expected height %d but received %d", tc.expectedHeight, result)
			}
		})
	}
}

func TestCreateBinarySubTreeStringsEdgeCases(t *testing.T) {
	var testCases = []struct {
		a              []string
		index          int
		expectedHeight int
	}{
		{a: nil, index: 0, expectedHeight: 0},
		{a: []string{}, index: 0, expectedHeight: 0},
		{a: []string{"nil"}, index: 0, expectedHeight: 0},
		{a: []string{"nil", "2"}, index: 0, expectedHeight: 0},
		{a: []string{"1"}, index: 2, expectedHeight: 0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.a), func(t *testing.T) {
			root := CreateBinarySubtreeStrings(tc.a, tc.index)
			if root != nil {
				t.Errorf("Expected tree root to be nil but received %d", root)
			}
			result := SubtreeHeight(root)
			if tc.expectedHeight != result {
				t.Errorf("Expected height %d but received %d", tc.expectedHeight, result)
			}
		})
	}
}

func TestCreateBinaryTree(t *testing.T) {
	for _, tc := range BinaryTreeTestCases {
		t.Run(fmt.Sprintf("%#v", tc.a), func(t *testing.T) {
			root := CreateBinaryTree(tc.a)
			tree := BinaryTree{Root: root}
			treeHeight := tree.Height()
			if tc.height != treeHeight {
				t.Errorf("Expected height %d but received %d", tc.height, treeHeight)
			}
		})
	}
}

func TestCreateBinarySubTreeSuccess(t *testing.T) {
	var testCases = []struct {
		a              []int
		index          int
		expectedHeight int
	}{
		{a: []int{1}, index: 0, expectedHeight: 1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.a), func(t *testing.T) {
			n := CreateBinarySubtree(tc.a, tc.index)
			result := SubtreeHeight(n)
			if tc.expectedHeight != result {
				t.Errorf("Expected height %d but received %d", tc.expectedHeight, result)
			}
		})
	}
}

func TestCreateBinarySubTreeEdgeCases(t *testing.T) {
	var testCases = []struct {
		a              []int
		index          int
		expectedHeight int
	}{
		{a: []int{}, index: 0, expectedHeight: 0},
		{a: []int{1, 2}, index: 1, expectedHeight: 1},
		{a: []int{1, 2}, index: 0, expectedHeight: 2},
		{a: []int{1, -1, 2, 3, 4}, index: 3, expectedHeight: 1},
		{a: []int{1, -1, 2, 3, 4}, index: 4, expectedHeight: 1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v with index %v", tc.a, tc.index), func(t *testing.T) {
			n := CreateBinarySubtree(tc.a, tc.index)
			if len(tc.a) == 0 {
				if n != nil {
					t.Fatal("Should have received a nil pointer when creating an empty subtree")
				}
			} else {
				result := SubtreeHeight(n)
				if tc.expectedHeight != result {
					t.Errorf("Expected height %d but received %d", tc.expectedHeight, result)
				}
			}
		})
	}
}

func TestPreOrderRecursive(t *testing.T) {
	for _, tc := range BinaryTreeTestCases {
		t.Run(fmt.Sprintf("PreOrderTraversal of %#v ", tc.a), func(t *testing.T) {
			result := TraversePreOrderRecursive(tc.tree.Root)
			if tc.preOrderTraversal != result {
				t.Errorf("Expected: %v but received %v", tc.preOrderTraversal, result)
			}
		})
	}
}

func TestPreOrderIterative(t *testing.T) {
	for _, tc := range BinaryTreeTestCases {
		t.Run(fmt.Sprintf("PreOrderTraversal of %v ", tc.a), func(t *testing.T) {
			result := TraversePreOrder(tc.tree.Root)
			if tc.preOrderTraversal != result {
				t.Errorf("Expected: %v but received %v", tc.preOrderTraversal, result)
			}
		})
	}
}

func TestInOrderTraversalIterative(t *testing.T) {
	for _, tc := range BinaryTreeTestCases {
		t.Run(fmt.Sprintf("%#v", tc.a), func(t *testing.T) {
			result := TraverseInOrder(tc.tree.Root)
			if tc.inOrderTraversal != result {
				t.Errorf("Expected: %v but received %v", tc.inOrderTraversal, result)
			}
		})
	}
}

func TestInOrderTraversalRecursive(t *testing.T) {
	for _, tc := range BinaryTreeTestCases {
		t.Run(fmt.Sprintf("%#v", tc.a), func(t *testing.T) {
			result := TraverseInOrderRecursive(tc.tree.Root)
			if tc.inOrderTraversal != result {
				t.Errorf("Expected: %#v but received %#v", tc.inOrderTraversal, result)
			}
		})
	}
}
