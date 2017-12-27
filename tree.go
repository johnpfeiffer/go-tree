package main

import "fmt"

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
