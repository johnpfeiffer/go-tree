package gotree

// BinarySearchTree https://en.wikipedia.org/wiki/Binary_search_tree
type BinarySearchTree struct {
	Root *Node
}

// Height is the longest distance from the root to a leaf in a binary tree
func (tree *BinarySearchTree) Height() int {
	if tree.Root == nil || (tree.Root.Left == nil && tree.Root.Right == nil) {
		return 0
	}
	return SubtreeHeight(tree.Root) - 1
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
			if current.Left == nil {
				current.Left = &Node{Data: target}
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &Node{Data: target}
				return
			}
			current = current.Right
		}
	}
}
