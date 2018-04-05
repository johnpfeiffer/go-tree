package gotree

import "strconv"

// BinaryTree https://en.wikipedia.org/wiki/Binary_tree
type BinaryTree struct {
	Root *Node
}

// Node contains data (and usually a value or a pointer to a value) and pointers to the child nodes
type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

// CreateBinaryTree creates a tree and returns the root given a slice of integers
func CreateBinaryTree(a []int) *Node {
	if len(a) == 0 {
		return nil
	}
	root := &Node{Data: a[0]}
	if len(a) > 1 {
		root.Left = CreateBinarySubtree(a, 1)
	}
	if len(a) > 2 {
		root.Right = CreateBinarySubtree(a, 2)
	}
	return root
}

// CreateBinarySubtree recursively creates a tree given a slice  of integers
func CreateBinarySubtree(a []int, index int) *Node {
	if len(a) == 0 {
		return nil
	}
	if index >= len(a) {
		return nil
	}
	n := &Node{Data: a[index]}
	if (index*2 + 1) < len(a) {
		n.Left = CreateBinarySubtree(a, index*2+1)
	}
	if (index*2 + 2) < len(a) {
		n.Right = CreateBinarySubtree(a, index*2+2)
	}
	return n
}

// CreateBinarySubtreeStrings returns the root after recursively creating a tree given a slice of strings that are either an integer or "nil"
func CreateBinarySubtreeStrings(a []string, index int) *Node {
	if len(a) == 0 {
		return nil
	}
	if index >= len(a) {
		return nil
	}
	if a[index] == "nil" {
		return nil
	}
	value, err := strconv.Atoi(a[index])
	if err != nil {
		// TODO: logging?
		return nil
	}
	n := &Node{Data: value}
	leftIndex := index*2 + 1
	if leftIndex < len(a) {
		n.Left = CreateBinarySubtreeStrings(a, leftIndex)
	}
	rightIndex := index*2 + 2
	if rightIndex < len(a) {
		n.Right = CreateBinarySubtreeStrings(a, rightIndex)
	}
	return n
}

// Height is the longest distance from the root to a leaf in a binary tree, effectively counting the number of edges
// https://www.cs.cmu.edu/~adamchik/15-121/lectures/Trees/trees.html
func (tree *BinaryTree) Height() int {
	if tree.Root == nil || (tree.Root.Left == nil && tree.Root.Right == nil) {
		return 0
	}
	return SubtreeHeight(tree.Root) - 1
}

// SubtreeHeight recursively calculates the largest distance from a node in a binary tree
func SubtreeHeight(n *Node) int {
	if n == nil {
		return 0
	}
	leftMax := 0
	rightMax := 0
	if n.Left == nil && n.Right == nil {
		return 1
	}
	if n.Left != nil {
		leftMax = SubtreeHeight(n.Left)
	}
	if n.Right != nil {
		rightMax = SubtreeHeight(n.Right)
	}
	if leftMax > rightMax {
		return leftMax + 1
	}
	return rightMax + 1
}
