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

// Display shows the node data and continues recursively (in pre-order)
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

// TODO: in-order and post-order display

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
	if tree.Root == nil {
		return
	}
	if tree.Root.Data == target {
		RemoveRoot(tree)
		return
	}
	current := tree.Root
	parent := current
	for {
		switch {
		case current == nil:
			return
		case current.Data == target:
			RemoveNode(current, parent)
			return
		case current.Data < target:
			parent = current
			current = current.right
		case current.Data > target:
			parent = current
			current = current.left
		}
	}
}

// RemoveRoot handles the special case of removing the root node
func RemoveRoot(tree *BinarySearchTree) {
	if tree.Root == nil {
		return
	}
	switch {
	case tree.Root.right == nil && tree.Root.left == nil:
		tree.Root = nil // if pointers then tree.Root.Data = nil to prevent memory leaks
	case tree.Root.right != nil:
		if tree.Root.right.left == nil { // simplest hoisting case, e.g. root to leaf: 1 2(root) 5 6 becomes 1 5(root) 6
			originalRootLeft := tree.Root.left
			tree.Root = tree.Root.right
			tree.Root.left = originalRootLeft
		} else { // more complex hoisting case, e.g. 1 2(root) 5 4 3 , need to find the right subtree left most, then fix the right subtree
			parent := tree.Root.right
			for current := tree.Root.right; current.left != nil; current = current.left {
				parent = current
			}
			originalRootLeft := tree.Root.left
			originalRootRight := tree.Root.right
			replacementRoot := parent.left

			tree.Root = replacementRoot
			parent.left = replacementRoot.right
			replacementRoot.left = originalRootLeft
			replacementRoot.right = originalRootRight
		}
		return

	case tree.Root.left != nil:
		if tree.Root.left.right == nil { // simplest hoisting case, e.g. root to leaf: 0 1 2(root) 5 becomes 0 1(root) 5
			originalRootRight := tree.Root.right
			tree.Root = tree.Root.left
			tree.Root.right = originalRootRight
		} else { // more complex hoisting case, e.g. 1 0 2(root) 5 , need to find the left subtree right most, then fix the left subtree
			parent := tree.Root.left
			for current := tree.Root.left; current.right != nil; current = current.right {
				parent = current
			}
			originalRootLeft := tree.Root.left
			originalRootRight := tree.Root.right
			replacementRoot := parent.right

			tree.Root = replacementRoot
			parent.right = replacementRoot.left
			replacementRoot.left = originalRootLeft
			replacementRoot.right = originalRootRight
		}
		return

	default:
		fmt.Println("ERROR should never reach here")
	}
}

// RemoveNode also handles the special case of removing a leaf node
func RemoveNode(node, parent *Node) {
	if node == nil || parent == nil {
		fmt.Println("ERROR should never reach here with node or parent as nil")
		return
	}
	switch {
	case parent.left == node:
		if node.left != nil {
			parent.left = node.left // hoist the remaining child, it is ok if we re-assign nil
			node.left = nil         // reminder that for pointers node.Data = nil prevents memory leaks
		} else {
			parent.left = node.right // simple logic as we do not care if we re-assign nil
			node.right = nil
		}
	case parent.right == node:
		if node.left != nil {
			parent.right = node.left
			node.left = nil
		} else {
			parent.right = node.right
			node.right = nil
		}
	default:
		fmt.Println("ERROR should never reach here with parent not matching the child node")
	}
}
