package main

import (
	"fmt"
	"testing"
)

func TestInsertSimple(t *testing.T) {
	expected := 42
	t.Run(fmt.Sprintf("Insert %#v into a BST", expected), func(t *testing.T) {
		tree := BinarySearchTree{}
		assertEmpty(t, tree)

		tree.InsertValue(42)
		current := tree.Root
		if current == nil {
			t.Error("Root Node for the BST is unexpectedly nil")
		}
		assertLeafNode(t, expected, tree.Root)

		expected2 := expected + 1
		tree.InsertValue(expected2)
		if current.left != nil {
			t.Error("Root Node left child for the BST is unexpectedly not nil")
		}
		if current.right == nil {
			t.Error("Root Node right child for the BST is unexpectedly nil")
		}
		current = current.right
		assertLeafNode(t, expected2, current)
	})
}

func assertEmpty(t *testing.T, tree BinarySearchTree) {
	t.Helper()
	if tree.Root != nil {
		t.Error("Root Node for the Binary Search Tree should still be nil")
	}
}

func assertLeafNode(t *testing.T, expectedData int, n *Node) {
	t.Helper()
	if expectedData != n.Data {
		t.Error("\nExpected:", expectedData, "\nReceived: ", n.Data)
	}
	if n.left != nil {
		t.Error("Left pointer should be nil")
	}
	if n.right != nil {
		t.Error("Right pointer should be nil")
	}
}
