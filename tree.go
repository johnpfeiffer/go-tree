package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// TreeNode contains data (and usually a value or a pointer to a value) and pointers to the child nodes
type TreeNode struct {
	children []*TreeNode
	Data     int
}

// Tree https://en.wikipedia.org/wiki/Tree_(data_structure)
type Tree struct {
	Root *TreeNode
}

// Display returns all of the nodes in a tree
func (tree *Tree) Display() {
	DisplayTreeNode(tree.Root)
	fmt.Println()
}

// DisplayTreeNode shows all of the nodes (data) in a tree
func DisplayTreeNode(n *TreeNode) {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Data)
	for _, child := range n.children {
		DisplayTreeNode(child)
	}
}

// Add inserts a node as a leaf
func (tree *Tree) Add(n *TreeNode) {
	current := tree.Root
	for {
		if len(current.children) == 0 {
			current.children = append(current.children, n)
			return
		}
		current = current.children[0]
	}
}

// AddValue is a helper to wrap creating a new node
func (tree *Tree) AddValue(n int) {
	tree.Add(&TreeNode{Data: n})
}

// BinaryNode contains data
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	Data  int
}

// createBinaryTree creates a tree given a slice
func createBinaryTree(a []int, index int, n *BinaryNode) {
	n = &BinaryNode{Data: a[index]}
	if (index*2 + 1) < len(a) {
		createBinaryTree(a, index*2+1, n.left)
	}
	if (index*2 + 2) < len(a) {
		createBinaryTree(a, index*2+2, n.right)
	}
	fmt.Println(n, n.left, n.right)
}

// BinaryPreOrder shows the node data (in pre-order) and continues recursively https://en.wikipedia.org/wiki/Tree_traversal#Pre-order
func BinaryPreOrder(n *BinaryNode) string {
	var s string
	var b bytes.Buffer
	if n == nil {
		return ""
	}
	b.WriteString(fmt.Sprintf("%d ", n.Data))
	if n.left != nil {
		s += BinaryPreOrder(n.left)
	}
	if n.right != nil {
		s += BinaryPreOrder(n.right)
	}
	return b.String() + s
}

// TraverseLevelOrderIntsRaw stores the empty leaf node nulls too
func TraverseLevelOrderIntsRaw(n *Node) []string {
	result := []string{}
	if n == nil {
		return result
	}
	var q NodeQueue
	q.enqueue(n)
	q.enqueue(nil)
	for q.length() > 0 {
		current := q.dequeue()
		if current == nil {
			result = append(result, "")
		} else {
			result = append(result, strconv.Itoa(current.Data))
			q.enqueue(current.left)
			q.enqueue(current.right)
			// if current.left != nil {
			// 	q.enqueue(current.left)
			// }
			// if current.right != nil {
			// 	q.enqueue(current.right)
			// }
		}
	}
	return result
}
