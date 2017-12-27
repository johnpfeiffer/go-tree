package main

import "fmt"

// Node contains data (and usually a value or a pointer to a value) and pointers to the child nodes
type Node struct {
	left  *Node
	right *Node
	Data  int
}

// BinarySearchTree https://en.wikipedia.org/wiki/Binary_search_tree
type BinarySearchTree struct {
	Root *Node
}

// Display shows all of the nodes in a tree
func (tree *BinarySearchTree) Display() {
	Display(tree.Root)
	fmt.Println()
}

// Display shows the node data and continues recursively
func Display(n *Node) {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Data)
	if n.left != nil {
		Display(n.left)
	}
	if n.right != nil {
		Display(n.right)
	}
}
