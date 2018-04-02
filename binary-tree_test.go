package gotree

import (
	"fmt"
	"testing"
)

// Manually building each tree variation once
var BinaryTreeTestCases = []struct {
	a                 []int
	tree              *BinaryTree
	preOrderTraversal string
	height            int
}{
	{a: nil, tree: &BinaryTree{}, preOrderTraversal: "", height: 0},
	{a: []int{}, tree: &BinaryTree{}, preOrderTraversal: "", height: 0},
	{a: []int{1}, tree: &BinaryTree{Root: &Node{Data: 1}}, preOrderTraversal: "1 ", height: 0},
	{a: []int{1, 2}, tree: &BinaryTree{Root: &Node{Data: 1, Left: &Node{Data: 2}}}, preOrderTraversal: "1 2 ", height: 1},
	// perfect tree /\
	{a: []int{1, 2, 3}, tree: &BinaryTree{Root: &Node{Data: 1, Left: &Node{Data: 2}, Right: &Node{Data: 3}}},
		preOrderTraversal: "1 2 3 ", height: 1},
	// TODO
	// degenerate trees  / / \ \
	//                  /  \ /  \

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
				t.Error("\nExpected data values:", tc.preOrderTraversal, "\nReceived: ", result)
			}
		})
	}
}

func TestPreOrderIterative(t *testing.T) {
	for _, tc := range BinaryTreeTestCases {
		t.Run(fmt.Sprintf("PreOrderTraversal of %v ", tc.a), func(t *testing.T) {
			result := TraversePreOrder(tc.tree.Root)
			if tc.preOrderTraversal != result {
				t.Error("\nExpected data values:", tc.preOrderTraversal, "\nReceived: ", result)
			}
		})
	}
}
