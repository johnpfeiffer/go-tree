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
	height            int
}{
	{dataValues: []int{}, tree: BinarySearchTree{}, preOrderTraversal: "", height: 0},
	{dataValues: []int{}, tree: BinarySearchTree{Root: nil}, preOrderTraversal: "", height: 0},
	{dataValues: []int{2}, tree: BinarySearchTree{Root: &Node{Data: 2}}, preOrderTraversal: "2 ", height: 0},
	{dataValues: []int{2, 1}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 1}}},
		preOrderTraversal: "2 1 ", height: 1},
	{dataValues: []int{2, 5}, tree: BinarySearchTree{Root: &Node{Data: 2, right: &Node{Data: 5}}},
		preOrderTraversal: "2 5 ", height: 1},
	// perfect tree /\
	{dataValues: []int{2, 1, 5}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 1}, right: &Node{Data: 5}}},
		preOrderTraversal: "2 1 5 ", height: 1},
	// degenerate trees  / / \ \
	//                  /  \ /  \
	{dataValues: []int{2, 1, -1}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 1, left: &Node{Data: -1}}}},
		preOrderTraversal: "2 1 -1 ", height: 2},
	{dataValues: []int{2, 0, 1}, tree: BinarySearchTree{Root: &Node{Data: 2, left: &Node{Data: 0, right: &Node{Data: 1}}}},
		preOrderTraversal: "2 0 1 ", height: 2},
	{dataValues: []int{2, 5, 4}, tree: BinarySearchTree{Root: &Node{Data: 2, right: &Node{Data: 5, left: &Node{Data: 4}}}},
		preOrderTraversal: "2 5 4 ", height: 2},
	{dataValues: []int{2, 5, 6}, tree: BinarySearchTree{Root: &Node{Data: 2, right: &Node{Data: 5, right: &Node{Data: 6}}}},
		preOrderTraversal: "2 5 6 ", height: 2},

	// a few examples of even more complex trees
	// left subtree variants, excluding the uninteresting
	{dataValues: []int{2, 1, 0, -1}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 1,
				left: &Node{Data: 0,
					left: &Node{Data: -1}}}}},
		preOrderTraversal: "2 1 0 -1 ", height: 3},
	{dataValues: []int{2, 1, -1, 0}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 1,
				left: &Node{Data: -1,
					right: &Node{Data: 0}}}}},
		preOrderTraversal: "2 1 -1 0 ", height: 3},
	{dataValues: []int{2, 0, 1, -1}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 0,
				right: &Node{Data: 1},
				left:  &Node{Data: -1}}}},
		preOrderTraversal: "2 0 -1 1 ", height: 2},
	{dataValues: []int{2, -1, 0, 1}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: -1,
				right: &Node{Data: 0,
					right: &Node{Data: 1}}}}},
		preOrderTraversal: "2 -1 0 1 ", height: 3},
	// // both subtrees variants
	{dataValues: []int{2, 1, -1, 5}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 1,
				left: &Node{Data: -1}},
			right: &Node{Data: 5}}},
		preOrderTraversal: "2 1 -1 5 ", height: 2},
	{dataValues: []int{2, -1, 1, 5}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: -1,
				right: &Node{Data: 1}},
			right: &Node{Data: 5}}},
		preOrderTraversal: "2 -1 1 5 ", height: 2},
	{dataValues: []int{2, 1, 5, 4}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 1},
			right: &Node{Data: 5,
				left: &Node{Data: 4}}}},
		preOrderTraversal: "2 1 5 4 ", height: 2},
	{dataValues: []int{2, 1, 5, 6}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			left: &Node{Data: 1},
			right: &Node{Data: 5,
				right: &Node{Data: 6}}}},
		preOrderTraversal: "2 1 5 6 ", height: 2},
	// right subtree variants, excluding the uninteresting {2, 4, 5, 6}
	{dataValues: []int{2, 4, 6, 5}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			right: &Node{Data: 4,
				right: &Node{Data: 6,
					left: &Node{Data: 5}}}}},
		preOrderTraversal: "2 4 6 5 ", height: 3},
	{dataValues: []int{2, 6, 4, 5}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			right: &Node{Data: 6,
				left: &Node{Data: 4,
					right: &Node{Data: 5}}}}},
		preOrderTraversal: "2 6 4 5 ", height: 3},
	{dataValues: []int{2, 6, 5, 4}, tree: BinarySearchTree{
		Root: &Node{Data: 2,
			right: &Node{Data: 6,
				left: &Node{Data: 5,
					left: &Node{Data: 4}}}}},
		preOrderTraversal: "2 6 5 4 ", height: 3},
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
		assertNumber(t, "height", 1, tree.Height())
	})
}

func TestPreOrderRecursive(t *testing.T) {
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("PreOrderTraversal of %v ", tc.dataValues), func(t *testing.T) {
			result := TraversePreOrderRecursive(tc.tree.Root)
			if tc.preOrderTraversal != result {
				t.Error("\nExpected data values:", tc.preOrderTraversal, "\nReceived: ", result)
			}
		})
	}
}

func TestPreOrderIterative(t *testing.T) {
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
			tree := createBST(tc.dataValues)
			treeString := TraverseInOrder(tree.Root)
			if expected != treeString {
				t.Error("\nTest ERROR: Expected tree (string):", expected, "\nReceived: ", treeString)
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

func TestHeight(t *testing.T) {
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Height of tree %v ", tc.preOrderTraversal), func(t *testing.T) {
			expected := sortedIntsString(tc.dataValues)
			tree := createBST(tc.dataValues)
			treeString := TraverseInOrder(tree.Root)
			if expected != treeString {
				t.Error("\nTest ERROR: Expected tree (string):", expected, "\nReceived: ", treeString)
			}
			assertNumber(t, "height", tc.height, tree.Height())
			assertNumber(t, "height", tc.height, tc.tree.Height())
		})
	}
}

func TestMinimumDepthDFS(t *testing.T) {
	var testCases = []struct {
		dataValues []int
		height     int
		minDepth   int
	}{
		// right sided trees
		{dataValues: []int{2}, height: 0, minDepth: 1},
		{dataValues: []int{2, 3}, height: 1, minDepth: 2},
		{dataValues: []int{2, 3, 5}, height: 2, minDepth: 3},
		{dataValues: []int{2, 3, 5, 1}, height: 2, minDepth: 2},
		{dataValues: []int{2, 3, 5, 1, 4}, height: 3, minDepth: 2},
		{dataValues: []int{2, 3, 5, 1, 4, 0}, height: 3, minDepth: 3},
		// left sided trees
		{dataValues: []int{2, -3}, height: 1, minDepth: 2},
		{dataValues: []int{2, -3, -2}, height: 2, minDepth: 3},
		{dataValues: []int{2, -3, -2}, height: 2, minDepth: 3},
		{dataValues: []int{2, -3, -2, -1}, height: 3, minDepth: 4},
		{dataValues: []int{2, -3, -2, -4}, height: 2, minDepth: 3},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("right sided tree %v", tc.dataValues), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			assertNumber(t, "height", tc.height, tree.Height())
			assertNumber(t, "minimum depth", tc.minDepth, tree.MinimumDepth())
		})
	}
}

func TestTraverseLevelOrder(t *testing.T) {
	var testCases = []struct {
		dataValues []int
		height     int
		expected   string
	}{
		// TODO: just extend the original BST testcases?
		// right sided trees
		{dataValues: []int{2}, height: 0, expected: "2"},
		{dataValues: []int{2, 3}, height: 1, expected: "2 3"},
		{dataValues: []int{2, 3, 5}, height: 2, expected: "2 3 5"},
		{dataValues: []int{2, 3, 5, 1}, height: 2, expected: "2 1 3 5"},
		{dataValues: []int{2, 3, 1, 5, 0}, height: 2, expected: "2 1 3 0 5"},
		{dataValues: []int{2, 3, 1, 0, 5, 4}, height: 3, expected: "2 1 3 0 5 4"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("right sided tree %v", tc.dataValues), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			assertNumber(t, "height", tc.height, tree.Height())
			result := TraverseLevelOrder(tree.Root)
			if result != tc.expected {
				t.Error("Expected", tc.expected, "but received", result)
			}
		})
	}
}

func TestFindNothing(t *testing.T) {
	nonExistentValue := 1001
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Should not find %v in tree %v ", nonExistentValue, tc.preOrderTraversal), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			result := tree.Find(nonExistentValue)
			result2 := tc.tree.Find(nonExistentValue)
			if nil != result || nil != result2 {
				t.Error("\nExpected to not find ", nonExistentValue, " in: ", tc.preOrderTraversal, " but got ", result, "or", result2)
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

func TestFindParentNil(t *testing.T) {
	nonExistentValue := 1001
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("Should not find %v in tree %v ", nonExistentValue, tc.dataValues), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			result := tree.Find(nonExistentValue)
			if nil != result {
				t.Error("\nExpected to not find ", nonExistentValue, " in: ", tc.dataValues, " but got ", result)
			}
		})
	}
}

// TestFindParentEasy leverages the fact that the first element inserted after the root become its child
func TestFindParentEasy(t *testing.T) {
	for _, tc := range BSTTestCases {
		t.Run(fmt.Sprintf("in tree %v", tc.dataValues), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			switch len(tc.dataValues) {
			case 0:
				result := FindParent(1001, tree.Root)
				if result != nil {
					t.Error("\nExpected nil when passed a nil node, but got", result)
				}
			case 1:
				result := FindParent(tc.dataValues[0], tree.Root)
				if tc.dataValues[0] != result.Data {
					t.Error("\nExpecting the parent is itself when searching for oneself, instead found", result)
				}
			default:
				result := FindParent(tc.dataValues[1], tree.Root)
				if tc.dataValues[0] != result.Data {
					t.Error("\nExpected parent data ", tc.dataValues[0], " in: ", tc.preOrderTraversal, " but got ", result)
				}
			}
		})
	}
}

func TestFindParentSuccess(t *testing.T) {
	var testCases = []struct {
		dataValues    []int
		target        int
		expectedData  int
		expectedLeft  *Node
		expectedRight *Node
	}{
		{dataValues: []int{2, 1, 3}, target: 3, expectedData: 2, expectedLeft: &Node{Data: 1}, expectedRight: &Node{Data: 3}},
		{dataValues: []int{1, 2, 3}, target: 3, expectedData: 2, expectedLeft: nil, expectedRight: &Node{Data: 3}},
		{dataValues: []int{2, 5, 3}, target: 3, expectedData: 5, expectedLeft: &Node{Data: 3}, expectedRight: nil},
		{dataValues: []int{2, 5, 3, 4}, target: 3, expectedData: 5, expectedLeft: &Node{Data: 3}, expectedRight: nil},
		{dataValues: []int{2, 5, 4, 3}, target: 3, expectedData: 4, expectedLeft: &Node{Data: 3}, expectedRight: nil},

		{dataValues: []int{2, -1, 0}, target: 0, expectedData: -1, expectedLeft: nil, expectedRight: &Node{Data: 0}},
		{dataValues: []int{2, -1, 0, 1}, target: 0, expectedData: -1, expectedLeft: nil, expectedRight: &Node{Data: 0}},
		{dataValues: []int{2, -1, 1, 0}, target: 0, expectedData: 1, expectedLeft: &Node{Data: 0}, expectedRight: nil},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Find parent of %v in tree %v ", tc.target, tc.dataValues), func(t *testing.T) {
			tree := createBST(tc.dataValues)
			result := FindParent(tc.target, tree.Root)
			if tc.expectedData != result.Data {
				t.Error("\nExpected parent data:", tc.expectedData, "\nReceived: ", result.Data)
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
			assertNumber(t, "height", tc.height, tree.Height())
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
	// TODO: struct that just extends the nested BSTTestCases
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
				assertNumber(t, "height", 0, tree.Height())
			case 3:
				if tc.dataValues[0] > tc.dataValues[1] && tc.dataValues[1] > tc.dataValues[2] {
					assertNumber(t, "height", 1, tree.Height())
				}
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

func TestReplaceLeft(t *testing.T) {
	var testCases = []struct {
		dataValues []int
		target     int
	}{
		// LEAF NODES
		{dataValues: []int{2, -1}, target: -1},
		{dataValues: []int{2, -1, 3}, target: -1},
		{dataValues: []int{2, 0, -1}, target: -1},
		{dataValues: []int{2, 0, -1, 1}, target: -1},
		// NO RIGHT SUBTREE
		{dataValues: []int{2, 1, 0}, target: 1},
		{dataValues: []int{2, 1, 0, -1}, target: 1},     // partial left subtree
		{dataValues: []int{2, 1, 0, -2, -1}, target: 1}, // perfect left subtree
		// NO LEFT SUBTREE
		{dataValues: []int{2, -2, 0}, target: -2},
		{dataValues: []int{2, -2, 0, -1}, target: -2}, // partial right subtree
		{dataValues: []int{2, -2, 0, -1}, target: -2}, // perfect right subtree
		// BOTH SUBTREES
		{dataValues: []int{2, -2, -6, 0}, target: -2},
		{dataValues: []int{2, -2, -6, 0, -7}, target: -2},
		{dataValues: []int{2, -2, -6, 0, -3}, target: -2},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Remove %v from tree %v ", tc.target, tc.dataValues), func(t *testing.T) {
			// GIVEN
			reducedDataValues := intRemoved(tc.target, tc.dataValues)
			expected := sortedIntsString(reducedDataValues)
			tree := createBST(tc.dataValues)
			// WHEN
			removeeParent := FindParent(tc.target, tree.Root)
			removeLeft(removeeParent)
			// THEN
			if expected != tree.Display() {
				t.Error("\nExpected tree (string):", expected, "\nReceived: ", tree.Display())
			}
		})
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	var testCases = []struct {
		dataValues []int
		v          int
		w          int
		expected   int
	}{
		// Defensive programming edge cases
		{dataValues: []int{}, v: -1, w: 1, expected: 0},      // tree of size 0
		{dataValues: []int{1}, v: 1, w: 1, expected: 1},      // tree of size 1
		{dataValues: []int{1}, v: -1, w: 1, expected: 1},     // not even in the tree?
		{dataValues: []int{0, 1}, v: 0, w: 1, expected: 0},   // tree of size 2?
		{dataValues: []int{0, -1}, v: -1, w: 0, expected: 0}, // tree of size 2 (left)?
		// Real but easy cases
		{dataValues: []int{0, 1, -1}, v: -1, w: 1, expected: 0}, // perfect tree
		{dataValues: []int{0, 1, -1}, v: 1, w: -1, expected: 0}, // perfect tree
		{dataValues: []int{0, 2, 1, 3}, v: 1, w: 3, expected: 2},
		{dataValues: []int{0, 2, 1, 3}, v: 3, w: 1, expected: 2},
		{dataValues: []int{0, 2, 1, 3, -1}, v: -1, w: 3, expected: 0},
		{dataValues: []int{0, 2, 1, 3, -1}, v: 3, w: -1, expected: 0},
		// perfect tree
		{dataValues: []int{0, 2, 1, 3, -2, -3, -1}, v: -1, w: 3, expected: 0},
		{dataValues: []int{0, 2, 1, 3, -2, -3, -1}, v: 3, w: -1, expected: 0},
		{dataValues: []int{0, 2, 1, 3, -2, -3, -1}, v: 3, w: -2, expected: 0},
		{dataValues: []int{0, 2, 1, 3, -2, -3, -1}, v: 3, w: -3, expected: 0},
		{dataValues: []int{0, 2, 1, 3, -2, -3, -1}, v: -3, w: 3, expected: 0},
		{dataValues: []int{0, 2, 1, 3, -2, -3, -1}, v: -3, w: 3, expected: 0},
		{dataValues: []int{0, 2, 1, 3, -2, -3, -1}, v: -2, w: -3, expected: -2}, // ancestor is the parent of the child?
		{dataValues: []int{0, 2, 1, 3, -2, -3, -1}, v: -3, w: -1, expected: -2},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("of %v and %v from BST %v should be %v",
			tc.v, tc.w, tc.expected, tc.dataValues), func(t *testing.T) {
			// GIVEN
			tree := createBST(tc.dataValues)
			// WHEN
			result := tree.LowestCommonAncestor(tc.v, tc.w)
			// THEN
			if len(tc.dataValues) == 0 {
				if result != nil {
					t.Errorf("Expected no result with an empty tree but received %#v", result)
				}
			} else {
				assertNumber(t, "Common Ancestor", tc.expected, result.Data)
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
		t.Error("Root Node for an empty Binary Search Tree should still be nil")
	}
	if tree.Height() != 0 {
		t.Error("Height for an empty Binary Search Tree should still be nil")
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
			t.Error("\nExpected", hint, "as nil but received", hint, "node with data:", result.Data)
			return
		}
	} else {
		if result == nil {
			t.Error("\nExpected", hint, "node with data: ", expected.Data, "but received nil")
			return
		}
		if expected.Data != result.Data {
			t.Error("\nExpected", hint, " data:", expected.Data, "\nReceived node with", hint, " data:", result.Data)
			return
		}
	}
}

func assertNumber(t *testing.T, hint string, expected, result int) {
	if expected != result {
		t.Error("\nExpected", hint, ":", expected, "\nReceived:", result)
	}
}
