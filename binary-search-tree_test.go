package main

import (
	"fmt"
	"testing"
)

// Manually building each tree variation once
var BSTTestCases = []struct {
	dataValues        []int
	tree              BinarySearchTree
	preOrderTraversal string
}{
	{dataValues: []int{}, tree: BinarySearchTree{}, preOrderTraversal: ""},
	{dataValues: []int{}, tree: BinarySearchTree{Root: nil}, preOrderTraversal: ""},
	{dataValues: []int{2}, tree: BinarySearchTree{Root: &Node{Data: 2}}, preOrderTraversal: "2 "},
	{dataValues: []int{2, 1}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 1}}}, preOrderTraversal: "2 1 "},
	{dataValues: []int{2, 5}, tree: BinarySearchTree{Root: &Node{Data: 2, right: &Node{Data: 5}}}, preOrderTraversal: "2 5 "},
	// perfect tree /\
	{dataValues: []int{2, 1, 5}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 1}, right: &Node{Data: 5}}},
		preOrderTraversal: "2 1 5 "},
	// degenerate trees  / / \ \
	//                  /  \ /  \
	{dataValues: []int{2, 1, -1}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 1, left: &Node{Data: -1}}}},
		preOrderTraversal: "2 1 -1 "},
	{dataValues: []int{2, 0, 1}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 0, right: &Node{Data: 1}}}},
		preOrderTraversal: "2 0 1 "},
	{dataValues: []int{2, 5, 4}, tree: BinarySearchTree{Root: &Node{Data: 2, right: &Node{Data: 5, left: &Node{Data: 4}}}},
		preOrderTraversal: "2 5 4 "},
	{dataValues: []int{2, 5, 6}, tree: BinarySearchTree{Root: &Node{Data: 2, right: &Node{Data: 5, right: &Node{Data: 6}}}},
		preOrderTraversal: "2 5 6 "},

	// a few examples of even more complex trees
	// left subtree variants, excluding the uninteresting
	{dataValues: []int{2, 1, 0, -1}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 1,
				left: &Node{Data: 0,
					left: &Node{Data: -1}}}}},
		preOrderTraversal: "2 1 0 -1 "},

	{dataValues: []int{2, 1, -1, 0}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 1,
				left: &Node{Data: -1,
					right: &Node{Data: 0}}}}},
		preOrderTraversal: "2 1 -1 0 "},
	{dataValues: []int{2, 0, 1, -1}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 0,
				left:  &Node{Data: -1},
				right: &Node{Data: 1}}}},
		preOrderTraversal: "2 0 -1 1 "},
	{dataValues: []int{2, -1, 0, 1}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: -1,
				right: &Node{Data: 0,
					right: &Node{Data: 1}}}}},
		preOrderTraversal: "2 -1 0 1 "},
	// both subtrees variants
	{dataValues: []int{2, 1, -1, 5}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 1,
				left: &Node{Data: -1}},
			right: &Node{Data: 5}}},
		preOrderTraversal: "2 1 -1 5 "},
	{dataValues: []int{2, -1, 1, 5}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: -1,
				right: &Node{Data: 1}},
			right: &Node{Data: 5}}},
		preOrderTraversal: "2 -1 1 5 "},
	{dataValues: []int{2, 1, 5, 4}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 1},
			right: &Node{Data: 5,
				left: &Node{Data: 4}}}},
		preOrderTraversal: "2 1 5 4 "},
	{dataValues: []int{2, 1, 5, 6}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 1},
			right: &Node{Data: 5,
				right: &Node{Data: 6}}}},
		preOrderTraversal: "2 1 5 6 "},
	// right subtree variants, excluding the uninteresting {2, 4, 5, 6}
	{dataValues: []int{2, 4, 6, 5}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			right: &Node{Data: 4,
				right: &Node{Data: 6,
					left: &Node{Data: 5}}}}},
		preOrderTraversal: "2 4 6 5 "},
	{dataValues: []int{2, 6, 4, 5}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			right: &Node{Data: 6,
				left: &Node{Data: 4,
					right: &Node{Data: 5}}}}},
		preOrderTraversal: "2 6 4 5 "},
	{dataValues: []int{2, 6, 5, 4}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			right: &Node{Data: 6,
				left: &Node{Data: 5,
					left: &Node{Data: 4}}}}},
		preOrderTraversal: "2 6 5 4 "},
}

func TestInsertSimple(t *testing.T) {
	expected := 42
	t.Run(fmt.Sprintf("Insert %#v into a BST", expected), func(t *testing.T) {
		tree := BinarySearchTree{}
		assertEmpty(t, tree)

		tree.InsertValue(42)
		current := tree.Root
		if current == nil {
			t.Error("Root Node for the BST is unexpectedly nil")
		}
		assertLeafNode(t, expected, tree.Root)

		expected2 := expected + 1
		tree.InsertValue(expected2)
		if current.left != nil {
			t.Error("Root Node left child for the BST is unexpectedly not nil")
		}
		if current.right == nil {
			t.Error("Root Node right child for the BST is unexpectedly nil")
		}
		current = current.right
		assertLeafNode(t, expected2, current)
	})
}

func TestPreOrder(t *testing.T) {
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("PreOrderTraversal of %v ", tc.dataValues), func(t *testing.T) {
			result := TraversePreOrder(tc.tree.Root)
			if tc.preOrderTraversal != result {
				t.Error("\nExpected data values:", tc.preOrderTraversal, "\nReceived: ", result)
			}
		})
	}
}

func TestInOrder(t *testing.T) {
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Traversal of %v ", tc.dataValues), func(t *testing.T) {
			expected := sortedIntsString(tc.dataValues)
			expectedBST := createBST(tc.dataValues)
			expectedBSTString := TraverseInOrder(expectedBST.Root)
			if expected != expectedBSTString {
				t.Error("\nTest ERROR: Expected tree (string):", expected, "\nReceived: ", expectedBSTString)
			}
			result := TraverseInOrder(tc.tree.Root)
			if expected != result {
				t.Error("\nExpected:", expected, "\nReceived: ", result)
			}
		})
	}
}

// func TestDisplay(t *testing.T) for tree.Display() is not very valuable since it is only a convenience wrapper

func TestInsertAdvanced(t *testing.T) {
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Insert to create tree %v ", tc.preOrderTraversal), func(t *testing.T) {
			expected := sortedIntsString(tc.dataValues)
			tree := createBST(tc.dataValues)
			result := tree.Display()
			if expected != result {
				t.Error("\nExpected inserted tree (string):", expected, "\nReceived (in-order): ", result)
			}
			preOrderResult := TraversePreOrder(tree.Root)
			if tc.preOrderTraversal != preOrderResult {
				t.Error("\nExpected inserted tree (string):", tc.preOrderTraversal, "\nReceived (pre-order): ", preOrderResult)
			}
		})
	}
}

func TestFindSuccess(t *testing.T) {
	var testCases = []struct {
		dataValues    []int
		target        int
		expectedLeft  *Node
		expectedRight *Node
	}{
		{dataValues: []int{2}, target: 2, expectedLeft: nil, expectedRight: nil},
		{dataValues: []int{2, 1}, target: 2, expectedLeft: &Node{Data: 1}, expectedRight: nil},
		{dataValues: []int{2, 1, 3}, target: 2, expectedLeft: &Node{Data: 1}, expectedRight: &Node{Data: 3}},
		{dataValues: []int{2, 1}, target: 1, expectedLeft: nil, expectedRight: nil},
		{dataValues: []int{2, 3}, target: 3, expectedLeft: nil, expectedRight: nil},

		{dataValues: []int{2, 1, -1}, target: 2, expectedLeft: &Node{Data: 1}, expectedRight: nil},
		{dataValues: []int{2, 1, -1}, target: 1, expectedLeft: &Node{Data: -1}, expectedRight: nil},
		{dataValues: []int{2, 0, 1}, target: 0, expectedLeft: nil, expectedRight: &Node{Data: 1}},
		{dataValues: []int{2, 1, -1}, target: -1, expectedLeft: nil, expectedRight: nil},
		{dataValues: []int{2, 0, 1}, target: 1, expectedLeft: nil, expectedRight: nil},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Find %v in tree %v ", tc.target, tc.dataValues), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			result := tree.Find(tc.target)
			if tc.target != result.Data {
				t.Error("\nExpected data:", tc.target, "\nReceived node with data: ", result.Data)
			}
			assertNode(t, tc.expectedLeft, result.left, "left")
			assertNode(t, tc.expectedRight, result.right, "right")
		})
	}
}

func TestFindRightMostParent(t *testing.T) {
	var testCases = []struct {
		dataValues    []int
		expectedData  int
		expectedLeft  *Node
		expectedRight *Node
	}{
		{dataValues: []int{2}, expectedData: 2, expectedLeft: nil, expectedRight: nil},
		{dataValues: []int{2, 1}, expectedData: 2, expectedLeft: &Node{Data: 1}, expectedRight: nil},
		{dataValues: []int{2, 1, 3}, expectedData: 2, expectedLeft: &Node{Data: 1}, expectedRight: &Node{Data: 3}},
		{dataValues: []int{2, 4, 3}, expectedData: 2, expectedLeft: nil, expectedRight: &Node{Data: 4}},
		{dataValues: []int{2, 5, 4, 3}, expectedData: 2, expectedLeft: nil, expectedRight: &Node{Data: 5}},
		{dataValues: []int{2, 5, 3, 4}, expectedData: 2, expectedLeft: nil, expectedRight: &Node{Data: 5}},

		{dataValues: []int{2, 3, 4}, expectedData: 3, expectedLeft: nil, expectedRight: &Node{Data: 4}},
		{dataValues: []int{2, 3, 5, 4}, expectedData: 3, expectedLeft: nil, expectedRight: &Node{Data: 5}},
		{dataValues: []int{2, 1, 0, 3, 4}, expectedData: 3, expectedLeft: nil, expectedRight: &Node{Data: 4}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("from tree %v ", tc.dataValues), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			result := FindRightMostParent(tree.Root)
			if tc.expectedData != result.Data {
				t.Error("\nExpected parent data:", tc.expectedData, "\nReceived: ", result.Data)
			}
			assertNode(t, tc.expectedLeft, result.left, "left")
			assertNode(t, tc.expectedRight, result.right, "right")
		})
	}
}

func TestFindLeftMostParent(t *testing.T) {
	var testCases = []struct {
		dataValues    []int
		expectedData  int
		expectedLeft  *Node
		expectedRight *Node
	}{
		{dataValues: []int{2}, expectedData: 2, expectedLeft: nil, expectedRight: nil},
		{dataValues: []int{2, 3}, expectedData: 2, expectedLeft: nil, expectedRight: &Node{Data: 3}},
		{dataValues: []int{2, 1, 3}, expectedData: 2, expectedLeft: &Node{Data: 1}, expectedRight: &Node{Data: 3}},
		{dataValues: []int{2, 0, 1}, expectedData: 2, expectedLeft: &Node{Data: 0}, expectedRight: nil},
		{dataValues: []int{2, -1, 0, 1}, expectedData: 2, expectedLeft: &Node{Data: -1}, expectedRight: nil},
		{dataValues: []int{2, -1, 1, 0}, expectedData: 2, expectedLeft: &Node{Data: -1}, expectedRight: nil},

		{dataValues: []int{2, 1, 0}, expectedData: 1, expectedLeft: &Node{Data: 0}, expectedRight: nil},
		{dataValues: []int{2, 1, -1, 0}, expectedData: 1, expectedLeft: &Node{Data: -1}, expectedRight: nil},
		{dataValues: []int{2, 0, -1, 1}, expectedData: 0, expectedLeft: &Node{Data: -1}, expectedRight: &Node{Data: 1}},
		{dataValues: []int{2, 0, 1, -1}, expectedData: 0, expectedLeft: &Node{Data: -1}, expectedRight: &Node{Data: 1}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("from tree %v ", tc.dataValues), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			result := FindLeftMostParent(tree.Root)
			if tc.expectedData != result.Data {
				t.Error("\nExpected parent data:", tc.expectedData, "\nReceived: ", result.Data)
			}
			assertNode(t, tc.expectedLeft, result.left, "left")
			assertNode(t, tc.expectedRight, result.right, "right")

		})
	}
}

func TestFindNothing(t *testing.T) {
	nonExistentValue := 1001
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Should not find %v in tree %v ", nonExistentValue, tc.preOrderTraversal), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			result := tree.Find(nonExistentValue)
			if nil != result {
				t.Error("\nExpected to not find ", nonExistentValue, " in: ", tc.preOrderTraversal, " but got ", result)
			}
		})
	}
}

func TestRemoveValueSimpleNonExistent(t *testing.T) {
	nonExistentValue := 1001
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Removing %v should have no effect on tree %v ", nonExistentValue, tc.preOrderTraversal), func(t *testing.T) {
			expected := sortedIntsString(tc.dataValues)
			tree := createBST(tc.dataValues)
			tree.RemoveValue(nonExistentValue)
			if len(tc.dataValues) == 0 {
				assertEmpty(t, tree)
			}
			if len(tc.dataValues) == 1 {
				assertLeafNode(t, tree.Root.Data, tree.Root)
			}
			result := tree.Display()
			if expected != result {
				t.Error("\nExpected same tree (string):", expected, "\nReceived: ", tree.Display())
			}
			preOrderResult := TraversePreOrder(tree.Root)
			if tc.preOrderTraversal != preOrderResult {
				t.Error("\nExpected inserted tree (string):", tc.preOrderTraversal, "\nReceived (pre-order): ", preOrderResult)
			}
		})
	}
}

/*
func TestRemoveValueAdvanced(t *testing.T) {
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Removing a value from tree %v ", tc.dataValues), func(t *testing.T) {
			for _, v := range tc.dataValues {
				reducedDataValues := intRemoved(v, tc.dataValues)
				expected := sortedIntsString(reducedDataValues)
				tree := createBST(tc.dataValues)
				tree.RemoveValue(v)
				if len(reducedDataValues) == 0 {
					assertEmpty(t, tree)
				}
				if len(reducedDataValues) == 1 {
					assertLeafNode(t, tree.Root.Data, tree.Root)
				}
				result := tree.Display()
				if expected != result {
					t.Error("\nExpected tree (string):", expected, "\nReceived: ", result)
				}
			}
		})
	}
}
*/

func TestRemoveValueSimpleSuccess(t *testing.T) {
	var testCases = []struct {
		dataValues         []int
		target             int
		expectedTree       BinarySearchTree
		expectedTreeString string
	}{
		{dataValues: []int{2}, target: 2, expectedTree: BinarySearchTree{Root: nil}, expectedTreeString: ""},
		{dataValues: []int{2, 5}, target: 2, expectedTree: BinarySearchTree{Root: &Node{Data: 5}}, expectedTreeString: "5 "},
		{dataValues: []int{2, 1}, target: 2, expectedTree: BinarySearchTree{Root: &Node{Data: 1}}, expectedTreeString: "1 "},
		{dataValues: []int{2, 5}, target: 5, expectedTree: BinarySearchTree{Root: &Node{Data: 2}}, expectedTreeString: "2 "},
		// {dataValues: []int{2, 1}, target: 1, expectedTree: BinarySearchTree{Root: &Node{Data: 2}}, expectedTreeString: "2 "},
		{dataValues: []int{2, 1, 5}, target: 2, expectedTree: BinarySearchTree{Root: &Node{Data: 5, left: &Node{Data: 1}}},
			expectedTreeString: "5 1 "},
		{dataValues: []int{2, 5, 1}, target: 2, expectedTree: BinarySearchTree{Root: &Node{Data: 5, left: &Node{Data: 1}}},
			expectedTreeString: "5 1 "},

		{dataValues: []int{2, 1, 0}, target: 2, expectedTree: BinarySearchTree{Root: &Node{Data: 1, left: &Node{Data: 0}}},
			expectedTreeString: "1 0 "},
		{dataValues: []int{2, 0, 1}, target: 2, expectedTree: BinarySearchTree{Root: &Node{Data: 1, left: &Node{Data: 0}}},
			expectedTreeString: "1 0 "},

		{dataValues: []int{2, 1, 5, 4}, target: 2, expectedTree: BinarySearchTree{
			Root: &Node{Data: 4, left: &Node{Data: 1}, right: &Node{Data: 5}}}, expectedTreeString: "4 1 5 "},
		{dataValues: []int{2, 1, 5, 4, 3}, target: 2, expectedTree: BinarySearchTree{
			Root: &Node{Data: 3, left: &Node{Data: 1}, right: &Node{Data: 5, left: &Node{Data: 4}}}}, expectedTreeString: "3 1 5 4 "},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Remove %v from tree %v ", tc.target, tc.dataValues), func(t *testing.T) {
			reducedDataValues := intRemoved(tc.target, tc.dataValues)
			expected := sortedIntsString(reducedDataValues)
			tree := createBST(tc.dataValues)
			tree.RemoveValue(tc.target)
			switch len(tc.dataValues) {
			case 0:
			case 1:
				assertEmpty(t, tree)
			case 2:
				assertLeafNode(t, tree.Root.Data, tree.Root)
			}
			// if expected != tc.expectedTreeString {
			// 	t.Error("\nTest ERROR: Expected tree (string):", tc.expectedTreeString, "\nReceived: ", tc.expectedTree.Display())
			// }
			if expected != tree.Display() {
				t.Error("\nExpected tree (string):", tc.expectedTreeString, "\nReceived: ", tree.Display())
			}
		})
	}
}

// HELPER FUNCTIONS

func intInSlice(target int, a []int) bool {
	for _, v := range a {
		if target == v {
			return true
		}
	}
	return false
}

func assertEmpty(t *testing.T, tree BinarySearchTree) {
	t.Helper()
	if tree.Root != nil {
		t.Error("Root Node for the Binary Search Tree should still be nil")
	}
}

func assertLeafNode(t *testing.T, expectedData int, n *Node) {
	t.Helper()
	if expectedData != n.Data {
		t.Error("\nExpected:", expectedData, "\nReceived: ", n.Data)
	}
	if n.left != nil {
		t.Error("Left pointer should be nil")
	}
	if n.right != nil {
		t.Error("Right pointer should be nil")
	}
}

func assertNode(t *testing.T, expected, result *Node, hint string) {
	if expected == nil {
		if result != nil {
			t.Error("\nExpected", hint, "as nil but received", hint, "node with data: ", result.Data)
			return
		}
	} else {
		if result == nil {
			t.Error("\nExpected", hint, "node with data: ", expected.Data, "but received nil")
			return
		}
		if expected.Data != result.Data {
			t.Error("\nExpected", hint, " data:", expected.Data, "\nReceived node with", hint, " data: ", result.Data)
			return
		}
	}
}
