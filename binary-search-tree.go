package main

import "fmt"
import "bytes"

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

// Display returns a string with the data from all of the nodes in the tree
func (tree *BinarySearchTree) Display() string {
	return Display(tree.Root)
}

// Display shows the node data and continues recursively
func Display(n *Node) string {
	var s string
	var b bytes.Buffer
	if n == nil {
		return ""
	}
	b.WriteString(fmt.Sprintf("%d ", n.Data))
	if n.left != nil {
		s += Display(n.left)
	}
	if n.right != nil {
		s += Display(n.right)
	}
	return b.String() + s
}

// Find returns the first node that has a matching key
func (tree *BinarySearchTree) Find(target int) *Node {
	current := tree.Root
	for {
		switch {
		case current == nil:
			return nil
		case current.Data == target:
			return current
		case current.Data > target:
			current = current.left
		case current.Data < target:
			current = current.right
		}
	}
}

// InsertValue adds data (with a new node) to the Binary Search Tree
func (tree *BinarySearchTree) InsertValue(target int) {
	if tree.Root == nil {
		tree.Root = &Node{Data: target}
		return
	}
	current := tree.Root
	for {
		if current.Data > target {
			if current.left == nil {
				current.left = &Node{Data: target}
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = &Node{Data: target}
				return
			}
			current = current.right
		}
	}
}

// RemoveValue removes the first node with the matching data
func (tree *BinarySearchTree) RemoveValue(target int) {
	current := tree.Root
	if current == nil {
		return
	}
	if tree.Root.Data == target {
		switch {
		case tree.Root.right == nil && tree.Root.left == nil:
			tree.Root = nil // if pointers then tree.Root.Data = nil to prevent memory leaks
		case tree.Root.right != nil:
			// Left Rotation may reduce the depth of the right subtree by one
			if tree.Root.right.left == nil {
				tree.Root = tree.Root.right
			}
		case tree.Root.left != nil:
			if tree.Root.left.right == nil {
				tree.Root = tree.Root.left
			}
		default:
			fmt.Println("ERROR should never reach here")
		}
	}
}
