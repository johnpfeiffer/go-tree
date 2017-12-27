package main

import "fmt"

// Node contains data (and usually a value or a pointer to a value) and pointers to the child nodes
type Node struct {
	children []*Node
	Data     int
}

// Tree https://en.wikipedia.org/wiki/Tree_(data_structure)
type Tree struct {
	Root *Node
}

// Display returns all of the nodes in a tree
func (tree *Tree) Display() {
	Display(tree.Root)
	fmt.Println()
}

// Display returns all of the nodes in a tree
func Display(n *Node) {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Data)
	for _, child := range n.children {
		Display(child)
	}
}

// Add inserts a node as a leaf
func (tree *Tree) Add(n *Node) {
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
	tree.Add(&Node{Data: n})
}
