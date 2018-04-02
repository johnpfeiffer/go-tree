package gotree

import (
	"bytes"
	"fmt"
)

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

// TraversePreOrderRecursive shows the node data (in pre-order) and continues recursively https://en.wikipedia.org/wiki/Tree_traversal#Pre-order
func TraversePreOrderRecursive(n *Node) string {
	var s string
	var b bytes.Buffer
	if n == nil {
		return ""
	}
	b.WriteString(fmt.Sprintf("%d ", n.Data))
	if n.Left != nil {
		s += TraversePreOrderRecursive(n.Left)
	}
	if n.Right != nil {
		s += TraversePreOrderRecursive(n.Right)
	}
	return b.String() + s
}

// TraversePreOrder shows the node data (in pre-order) and continues iteratively
func TraversePreOrder(n *Node) string {
	var b bytes.Buffer
	if n == nil {
		return ""
	}
	stack := []*Node{n}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		b.WriteString(fmt.Sprintf("%d ", current.Data))
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	return b.String()
}
