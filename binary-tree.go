package gotree

import "strconv"

// BinaryTree https://en.wikipedia.org/wiki/Binary_tree
type BinaryTree struct {
	Root *Node
}

// TODO: interfaces to interact with BinarySearchTrees

// Node contains data (and usually a value or a pointer to a value) and pointers to the child nodes
type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

// CreateBinarySubtree returns the root after recursively creating a tree given a slice of strings that are either an integer or "nil"
func CreateBinarySubtree(a []string, index int) *Node {
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
		n.Left = CreateBinarySubtree(a, leftIndex)
	}
	rightIndex := index*2 + 2
	if rightIndex < len(a) {
		n.Right = CreateBinarySubtree(a, rightIndex)
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

// MinimumDepth is a convenience wrapper for the number of nodes on the shortest path from a tree root to a leaf
func (tree *BinaryTree) MinimumDepth() int {
	if tree.Root == nil {
		return 0
	}
	return subtreeMinimumDepth(tree.Root, 1)
}

// subtreeMinimumDepth is a depth first recursive algorithm to find the shortest path from root to leaf
// the leaf node triggers the return
// every intermediate node needs to add one
func subtreeMinimumDepth(n *Node, depth int) int {
	leftMax := 0
	rightMax := 0
	if n.Left == nil && n.Right == nil {
		return depth
	}
	if n.Right == nil {
		return subtreeMinimumDepth(n.Left, depth+1)
	}
	if n.Left == nil {
		return subtreeMinimumDepth(n.Right, depth+1)
	}
	leftMax = subtreeMinimumDepth(n.Left, depth+1)
	rightMax = subtreeMinimumDepth(n.Right, depth+1)
	if leftMax < rightMax {
		return leftMax
	}
	return rightMax
}
