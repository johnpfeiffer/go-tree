package main

import (
	"fmt"
	"testing"
)

// Manually building each tree variation once
var BSTTestCases = []struct {
	dataValues []int
	tree       BinarySearchTree
	treeString string
}{
	{dataValues: []int{}, tree: BinarySearchTree{}, treeString: ""},
	{dataValues: []int{}, tree: BinarySearchTree{Root: nil}, treeString: ""},
	{dataValues: []int{2}, tree: BinarySearchTree{Root: &Node{Data: 2}}, treeString: "2 "},
	{dataValues: []int{2, 1}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 1}}}, treeString: "2 1 "},
	{dataValues: []int{2, 5}, tree: BinarySearchTree{Root: &Node{Data: 2, right: &Node{Data: 5}}}, treeString: "2 5 "},
	// perfect tree /\
	{dataValues: []int{2, 1, 5}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 1}, right: &Node{Data: 5}}},
		treeString: "2 1 5 "},
	// degenerate trees  / / \ \
	//                  /  \ /  \
	{dataValues: []int{2, 1, -1}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 1, left: &Node{Data: -1}}}},
		treeString: "2 1 -1 "},
	{dataValues: []int{2, 0, 1}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 0, right: &Node{Data: 1}}}},
		treeString: "2 0 1 "},
	{dataValues: []int{2, 5, 4}, tree: BinarySearchTree{Root: &Node{Data: 2, right: &Node{Data: 5, left: &Node{Data: 4}}}},
		treeString: "2 5 4 "},
	{dataValues: []int{2, 5, 6}, tree: BinarySearchTree{Root: &Node{Data: 2, right: &Node{Data: 5, right: &Node{Data: 6}}}},
		treeString: "2 5 6 "},
	// a few examples of even more complex trees
	{dataValues: []int{2, 1, -1, 0}, tree: BinarySearchTree{Root: &Node{
		Data: 2,
		left: &Node{Data: 1, left: &Node{Data: -1}, right: &Node{Data: 0}}}},
		treeString: "2 1 -1 0 "},
	{dataValues: []int{2, 1, -1, 5}, tree: BinarySearchTree{Root: &Node{
		Data:  2,
		left:  &Node{Data: 1, left: &Node{Data: -1}},
		right: &Node{Data: 5}}},
		treeString: "2 1 -1 5 "},
	{dataValues: []int{2, 1, 0, 5}, tree: BinarySearchTree{Root: &Node{
		Data:  2,
		left:  &Node{Data: 1, right: &Node{Data: 0}},
		right: &Node{Data: 5}}},
		treeString: "2 1 0 5 "},
	{dataValues: []int{2, 1, 5, 4}, tree: BinarySearchTree{Root: &Node{
		Data:  2,
		left:  &Node{Data: 1},
		right: &Node{Data: 5, left: &Node{Data: 4}}}},
		treeString: "2 1 5 4 "},
	{dataValues: []int{2, 5, 4, 6}, tree: BinarySearchTree{Root: &Node{
		Data:  2,
		right: &Node{Data: 5, left: &Node{Data: 4, right: &Node{Data: 6}}}}},
		treeString: "2 5 4 6 "},
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

func TestDisplay(t *testing.T) {
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Display node %v ", tc.treeString), func(t *testing.T) {
			result := tc.tree.Display()
			if tc.treeString != result {
				t.Error("\nExpected display string:", tc.treeString, "\nReceived: ", result)
			}
		})
	}
}

func TestInsertAdvanced(t *testing.T) {
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Insert to create tree %v ", tc.treeString), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			result := tree.Display()
			if tc.treeString != result {
				t.Error("\nExpected inserted tree (string):", tc.treeString, "\nReceived: ", result)
			}
		})
	}
}

func TestFindSuccess(t *testing.T) {
	var testCases = []struct {
		dataValues         []int
		target             int
		expectedLeftChild  *Node
		expectedRightChild *Node
	}{
		{dataValues: []int{2}, target: 2, expectedLeftChild: nil, expectedRightChild: nil},
		{dataValues: []int{2, 1}, target: 2, expectedLeftChild: &Node{Data: 1}, expectedRightChild: nil},
		{dataValues: []int{2, 1, 3}, target: 2, expectedLeftChild: &Node{Data: 1}, expectedRightChild: &Node{Data: 3}},
		{dataValues: []int{2, 1}, target: 1, expectedLeftChild: nil, expectedRightChild: nil},
		{dataValues: []int{2, 3}, target: 3, expectedLeftChild: nil, expectedRightChild: nil},

		{dataValues: []int{2, 1, -1}, target: 2, expectedLeftChild: &Node{Data: 1}, expectedRightChild: nil},
		{dataValues: []int{2, 1, -1}, target: 1, expectedLeftChild: &Node{Data: -1}, expectedRightChild: nil},
		{dataValues: []int{2, 0, 1}, target: 0, expectedLeftChild: nil, expectedRightChild: &Node{Data: 1}},
		{dataValues: []int{2, 1, -1}, target: -1, expectedLeftChild: nil, expectedRightChild: nil},
		{dataValues: []int{2, 0, 1}, target: 1, expectedLeftChild: nil, expectedRightChild: nil},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Find %v in tree %v ", tc.target, tc.dataValues), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			result := tree.Find(tc.target)
			if tc.target != result.Data {
				t.Error("\nExpected data:", tc.target, "\nReceived node with data: ", result.Data)
			}
			if tc.expectedLeftChild == nil {
				if result.left != nil {
					t.Error("\nExpected left child nil but received node with left child data: ", result.left.Data)
				}
			} else {
				if tc.expectedLeftChild.Data != result.left.Data {
					t.Error("\nExpected left child data:", tc.expectedLeftChild.Data, "\nReceived node with left child data: ", result.left.Data)
				}
			}
			if tc.expectedRightChild == nil {
				if result.right != nil {
					t.Error("\nExpected right child nil but received node with right child data: ", result.right.Data)
				}
			} else {
				if tc.expectedRightChild.Data != result.right.Data {
					t.Error("\nExpected right child data:", tc.expectedRightChild.Data, "\nReceived node with right child data: ", result.right.Data)
				}
			}
		})
	}
}

func TestFindNothing(t *testing.T) {
	nonExistentValue := 1001
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Should not find %v in tree %v ", nonExistentValue, tc.treeString), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			result := tree.Find(nonExistentValue)
			if nil != result {
				t.Error("\nExpected to not find ", nonExistentValue, " in: ", tc.treeString, " but got ", result)
			}
		})
	}
}

func TestRemoveValueSimpleNonExistent(t *testing.T) {
	nonExistentValue := 1001
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Removing %v should have no effect on tree %v ", nonExistentValue, tc.treeString), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			tree.RemoveValue(nonExistentValue)
			if len(tc.dataValues) == 0 {
				assertEmpty(t, tree)
			}
			if len(tc.dataValues) == 1 {
				assertLeafNode(t, tree.Root.Data, tree.Root)
			}
			if tc.treeString != tree.Display() {
				t.Error("\nExpected same tree (string):", tc.treeString, "\nReceived: ", tree.Display())
			}
		})
	}

}

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
		{dataValues: []int{2, 1}, target: 1, expectedTree: BinarySearchTree{Root: &Node{Data: 2}}, expectedTreeString: "2 "},
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
			tree := createBST(tc.dataValues)
			tree.RemoveValue(tc.target)
			switch len(tc.dataValues) {
			case 0:
			case 1:
				assertEmpty(t, tree)
			case 2:
				assertLeafNode(t, tree.Root.Data, tree.Root)
			}
			if tc.expectedTree.Display() != tc.expectedTreeString {
				t.Error("\nTest ERROR: Expected tree (string):", tc.expectedTreeString, "\nReceived: ", tc.expectedTree.Display())
			}
			if tc.expectedTreeString != tree.Display() {
				t.Error("\nExpected tree (string):", tc.expectedTreeString, "\nReceived: ", tree.Display())
			}
		})
	}
}

// HELPER FUNCTIONS
func createBST(a []int) BinarySearchTree {
	bst := BinarySearchTree{}
	for _, v := range a {
		bst.InsertValue(v)
	}
	return bst
}

func intInSlice(target int, a []int) bool {
	for _, v := range a {
		if target == v {
			return true
		}
	}
	return false
}

func intRemoved(target int, a []int) []int {
	count := 0 // only remove the first instance
	var result []int
	for i := 0; i < len(a)-1; i++ {
		if count == 0 && a[i] == target {
			result = append(result, a[i])
			count++
		}
	}
	return result
}

// TODO: Tree Comparison

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
