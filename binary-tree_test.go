package gotree

import (
	"fmt"
	"strconv"
	"testing"
)

// Manually building each tree variation once

var BinaryTreeEdgeTestCases = []struct {
	a                 []string
	tree              *BinaryTree
	preOrderTraversal string
	inOrderTraversal  string
	height            int
}{
	{a: nil, tree: &BinaryTree{}, preOrderTraversal: "", inOrderTraversal: "", height: 0},
	{a: []string{}, tree: &BinaryTree{}, preOrderTraversal: "", inOrderTraversal: "", height: 0},
}

var BinaryTreeTestCases = []struct {
	a                 []string
	tree              *BinaryTree
	preOrderTraversal string
	inOrderTraversal  string
	height            int
	minDepth          int
}{
	{a: []string{"1"}, tree: &BinaryTree{Root: &Node{Data: 1}}, preOrderTraversal: "1 ", inOrderTraversal: "1", height: 0, minDepth: 1},
	{a: []string{"1", "nil"}, tree: &BinaryTree{Root: &Node{Data: 1}}, preOrderTraversal: "1 ", inOrderTraversal: "1", height: 0, minDepth: 1},
	{a: []string{"1", "2"}, tree: &BinaryTree{Root: &Node{Data: 1, Left: &Node{Data: 2}}},
		preOrderTraversal: "1 2 ", inOrderTraversal: "2 1", height: 1, minDepth: 2},
	{a: []string{"1", "2", "nil", "nil"}, tree: &BinaryTree{Root: &Node{Data: 1, Left: &Node{Data: 2}}},
		preOrderTraversal: "1 2 ", inOrderTraversal: "2 1", height: 1, minDepth: 2},
	{a: []string{"1", "nil", "3"}, tree: &BinaryTree{Root: &Node{Data: 1, Left: nil, Right: &Node{Data: 3}}},
		preOrderTraversal: "1 3 ", inOrderTraversal: "1 3", height: 1, minDepth: 2},

	// perfect tree    1
	//                2 3
	{a: []string{"1", "2", "3"}, tree: &BinaryTree{Root: &Node{Data: 1, Left: &Node{Data: 2}, Right: &Node{Data: 3}}},
		preOrderTraversal: "1 2 3 ", inOrderTraversal: "2 1 3", height: 1, minDepth: 2},

	// degenerate trees  / / \ \
	//                  /  \ /  \
	{a: []string{"1", "2", "nil", "4"},
		tree:              &BinaryTree{Root: &Node{Data: 1, Left: &Node{Data: 2, Left: &Node{Data: 4}}, Right: nil}},
		preOrderTraversal: "1 2 4 ", inOrderTraversal: "4 2 1", height: 2, minDepth: 3},
	{a: []string{"1", "2", "nil", "nil", "5"},
		tree:              &BinaryTree{Root: &Node{Data: 1, Left: &Node{Data: 2, Right: &Node{Data: 5}}}},
		preOrderTraversal: "1 2 5 ", inOrderTraversal: "2 5 1", height: 2, minDepth: 3},
	{a: []string{"1", "nil", "3", "nil", "nil", "6"},
		tree:              &BinaryTree{Root: &Node{Data: 1, Right: &Node{Data: 3, Left: &Node{Data: 6}}}},
		preOrderTraversal: "1 3 6 ", inOrderTraversal: "1 6 3", height: 2, minDepth: 3},
	{a: []string{"1", "nil", "3", "nil", "nil", "nil", "7"},
		tree:              &BinaryTree{Root: &Node{Data: 1, Right: &Node{Data: 3, Right: &Node{Data: 7}}}},
		preOrderTraversal: "1 3 7 ", inOrderTraversal: "1 3 7", height: 2, minDepth: 3},

	// larger trees   1
	// 				2   3
	//			   4 5 6 7
	{a: []string{"1", "2", "3", "4"}, tree: &BinaryTree{Root: &Node{
		Data: 1, Left: &Node{Data: 2, Left: &Node{Data: 4}}, Right: &Node{Data: 3}}},
		preOrderTraversal: "1 2 4 3 ", inOrderTraversal: "4 2 1 3", height: 2, minDepth: 2},
	{a: []string{"1", "2", "3", "5"}, tree: &BinaryTree{Root: &Node{
		Data: 1, Left: &Node{Data: 2, Right: &Node{Data: 5}}, Right: &Node{Data: 3}}},
		preOrderTraversal: "1 2 5 3 ", inOrderTraversal: "2 5 1 3", height: 2, minDepth: 2},
	{a: []string{"1", "2", "3", "6"}, tree: &BinaryTree{Root: &Node{
		Data: 1, Left: &Node{Data: 2}, Right: &Node{Data: 3, Left: &Node{Data: 6}}}},
		preOrderTraversal: "1 2 3 6 ", inOrderTraversal: "2 1 6 3", height: 2, minDepth: 2},

	{a: []string{"1", "2", "3", "4", "5"}, tree: &BinaryTree{Root: &Node{
		Data: 1, Left: &Node{Data: 2, Left: &Node{Data: 4}, Right: &Node{Data: 5}}, Right: &Node{Data: 3}}},
		preOrderTraversal: "1 2 4 5 3 ", inOrderTraversal: "4 2 5 1 3", height: 2, minDepth: 2},
	{a: []string{"1", "2", "3", "4", "nil", "6"}, tree: &BinaryTree{Root: &Node{
		Data: 1, Left: &Node{Data: 2, Left: &Node{Data: 4}}, Right: &Node{Data: 3, Left: &Node{Data: 6}}}},
		preOrderTraversal: "1 2 4 3 6 ", inOrderTraversal: "4 2 1 6 3", height: 2, minDepth: 3},
	{a: []string{"1", "2", "3", "4", "nil", "nil", "7"}, tree: &BinaryTree{Root: &Node{
		Data: 1, Left: &Node{Data: 2, Left: &Node{Data: 4}}, Right: &Node{Data: 3, Right: &Node{Data: 7}}}},
		preOrderTraversal: "1 2 4 3 7 ", inOrderTraversal: "4 2 1 3 7", height: 2, minDepth: 3},
	{a: []string{"1", "2", "3", "nil", "nil", "6", "7"}, tree: &BinaryTree{Root: &Node{
		Data: 1, Left: &Node{Data: 2}, Right: &Node{Data: 3, Left: &Node{Data: 6}, Right: &Node{Data: 7}}}},
		preOrderTraversal: "1 2 3 6 7 ", inOrderTraversal: "2 1 6 3 7", height: 2, minDepth: 2},
}

func TestCreateBinarySubTreeSuccess(t *testing.T) {
	for _, tc := range BinaryTreeTestCases {
		t.Run(fmt.Sprintf("%v", tc.a), func(t *testing.T) {
			root := CreateBinarySubtree(tc.a, 0)
			expectedRootData, _ := strconv.Atoi(tc.a[0])
			if expectedRootData != root.Data {
				t.Errorf("Expected root value: %d but received %d", expectedRootData, root.Data)
			}
			tree := BinaryTree{Root: root}
			treeHeight := tree.Height()
			if tc.height != treeHeight {
				t.Errorf("Expected height %d but received %d", tc.height, treeHeight)
			}
			// TODO: use compareTree()
		})
	}
}

func TestCreateBinarySubTreeEdgeCases(t *testing.T) {
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
			root := CreateBinarySubtree(tc.a, tc.index)
			if root != nil {
				t.Errorf("Expected tree root to be nil but received %d", root)
			}
			result := SubtreeHeight(root)
			if tc.expectedHeight != result {
				t.Errorf("Expected height %d but received %d", tc.expectedHeight, result)
			}
			tree := BinaryTree{Root: root}
			assertEmptyTree(t, tree)
		})
	}
}

func TestPreOrderRecursive(t *testing.T) {
	for _, tc := range BinaryTreeTestCases {
		t.Run(fmt.Sprintf("PreOrderTraversal of %#v ", tc.a), func(t *testing.T) {
			result := TraversePreOrderRecursive(tc.tree.Root)
			if tc.preOrderTraversal != result {
				t.Errorf("Expected: %#v but received %#v", tc.preOrderTraversal, result)
			}
		})
	}
}

func TestPreOrderIterative(t *testing.T) {
	for _, tc := range BinaryTreeTestCases {
		t.Run(fmt.Sprintf("PreOrderTraversal of %v ", tc.a), func(t *testing.T) {
			result := TraversePreOrder(tc.tree.Root)
			if tc.preOrderTraversal != result {
				t.Errorf("Expected: %#v but received %#v", tc.preOrderTraversal, result)
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

func TestLevelOrder(t *testing.T) {
	for _, tc := range BinaryTreeTestCases {
		t.Run(fmt.Sprintf("%#v", tc.a), func(t *testing.T) {
			// takes advantage of the fact that the test tree data is always setup in a specific pattern
			intsOnly := removeNils(tc.a)
			expected := sortedIntsString(intsOnly)
			result := TraverseLevelOrder(tc.tree.Root)
			if expected != result {
				t.Errorf("Expected: %#v but received %#v", tc.inOrderTraversal, result)
			}
		})
	}
}

func TestMinimumDepthDFS(t *testing.T) {
	for _, tc := range BinaryTreeTestCases {
		t.Run(fmt.Sprintf("tree values: %v", tc.a), func(t *testing.T) {
			assertNumber(t, "height", tc.height, tc.tree.Height())
			assertNumber(t, "minimum depth", tc.minDepth, tc.tree.MinimumDepth())
		})
	}
}
