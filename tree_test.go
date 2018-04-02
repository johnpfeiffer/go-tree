package gotree

import (
	"fmt"
	"testing"
)

func TestTreeString(t *testing.T) {
	var testCases = []struct {
		node     *TreeNode
		expected string
	}{
		{node: nil, expected: ""},
		{node: &TreeNode{}, expected: "0"},
		{node: &TreeNode{Data: 42}, expected: "42"},
		{node: &TreeNode{Data: 42, Children: nil}, expected: "42"},
		{node: &TreeNode{Data: 42, Children: []*TreeNode{}}, expected: "42"},
		{node: &TreeNode{Data: 42, Children: []*TreeNode{&TreeNode{Data: 43}}}, expected: "42 43"},
		{node: &TreeNode{Data: 42,
			Children: []*TreeNode{&TreeNode{Data: 43}, &TreeNode{Data: 44}, &TreeNode{Data: 45}}},
			expected: "42 43 44 45"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.node), func(t *testing.T) {
			result := tc.node.String()
			if tc.expected != result {
				t.Errorf("Expected %s but received %s", tc.expected, result)
			}
		})
	}
	tree := &Tree{
		Root: &TreeNode{Data: 42,
			Children: []*TreeNode{
				&TreeNode{Data: 43},
				&TreeNode{Data: 44, Children: []*TreeNode{&TreeNode{Data: 100}}},
				&TreeNode{Data: 45},
			},
		},
	}
	expected := "44 100"
	t.Run(fmt.Sprintf("tree %v to %v", tree, expected), func(t *testing.T) {
		target := tree.Root.Children[1]
		result := target.String()
		if expected != result {
			t.Errorf("Expected %s but received %s", expected, result)
		}
	})
}

func TestTreeAddSuccess(t *testing.T) {
	// t.Skip()
	var testCases = []struct {
		tree     *Tree
		n        *TreeNode
		expected string
	}{
		{tree: &Tree{}, n: &TreeNode{Data: 42}, expected: "42"},
		{tree: &Tree{}, n: nil, expected: ""},
		{tree: &Tree{&TreeNode{Data: 42}}, n: &TreeNode{Data: 43}, expected: "42 43"},
		{tree: &Tree{&TreeNode{Data: 42}}, n: nil, expected: "42 "},
		{tree: &Tree{&TreeNode{Data: 42, Children: nil}}, n: &TreeNode{Data: 43}, expected: "42 43"},
		{tree: &Tree{&TreeNode{Data: 42, Children: []*TreeNode{}}}, n: &TreeNode{Data: 43}, expected: "42 43"},
		{tree: &Tree{&TreeNode{Data: 42, Children: []*TreeNode{&TreeNode{Data: 43}}}}, n: &TreeNode{Data: 100}, expected: "42 43 100"},
		{tree: &Tree{&TreeNode{Data: 42,
			Children: []*TreeNode{&TreeNode{Data: 43}, &TreeNode{Data: 44}}}}, n: &TreeNode{Data: 100}, expected: "42 43 100 44"},
		{tree: &Tree{&TreeNode{Data: 42,
			Children: []*TreeNode{&TreeNode{Data: 43, Children: []*TreeNode{&TreeNode{Data: 100}}}, &TreeNode{Data: 44}}}},
			n: &TreeNode{Data: 200}, expected: "42 43 100 200 44"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v Adds %v", tc.tree, tc.n), func(t *testing.T) {
			tc.tree.Add(tc.n)
			result := tc.tree.Root.String()
			if tc.expected != result {
				t.Errorf("Expected %s but received %s", tc.expected, result)
			}
		})
	}
}

func TestTreeAddError(t *testing.T) {
	var testCases = []struct {
		tree     *Tree
		expected string
	}{
		// test matrix of one to somehow magically get a nil value to call method functions because nil.Add() does not work
		{tree: nil, expected: "Cannot Add nodes to a nil pointer"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Nil tree pointer"), func(t *testing.T) {
			err := tc.tree.Add(&TreeNode{Data: 42})
			if err == nil {
				t.Errorf("Expected but did not receive an error when adding a leaf node to a nil pointer tree")
			}
			if err.Error() != tc.expected {
				t.Errorf("Expected the error: '%s' but received '%s'", tc.expected, err.Error())
			}
		})
		t.Run(fmt.Sprintf("AddValue nil tree pointer"), func(t *testing.T) {
			err := tc.tree.AddValue(42)
			if err == nil {
				t.Errorf("Expected but did not receive an error when adding a leaf node to a nil pointer tree")
			}
			if err.Error() != tc.expected {
				t.Errorf("Expected the error: '%s' but received '%s'", tc.expected, err.Error())
			}
		})
	}
}

func TestTreeAddValue(t *testing.T) {
	// t.Skip()
	var testCases = []struct {
		tree     *Tree
		n        int
		expected string
	}{
		{tree: &Tree{}, n: 42, expected: "42"},
		{tree: &Tree{&TreeNode{Data: 42}}, n: 43, expected: "42 43"},
		{tree: &Tree{&TreeNode{Data: 42, Children: nil}}, n: 43, expected: "42 43"},
		{tree: &Tree{&TreeNode{Data: 42, Children: []*TreeNode{}}}, n: 43, expected: "42 43"},
		{tree: &Tree{&TreeNode{Data: 42, Children: []*TreeNode{&TreeNode{Data: 43}}}}, n: 100, expected: "42 43 100"},
		{tree: &Tree{&TreeNode{Data: 42,
			Children: []*TreeNode{&TreeNode{Data: 43}, &TreeNode{Data: 44}}}}, n: 100, expected: "42 43 100 44"},
		{tree: &Tree{&TreeNode{Data: 42,
			Children: []*TreeNode{&TreeNode{Data: 43, Children: []*TreeNode{&TreeNode{Data: 100}}}, &TreeNode{Data: 44}}}},
			n: 200, expected: "42 43 100 200 44"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v Adds Value %v", tc.tree, tc.n), func(t *testing.T) {
			tc.tree.AddValue(tc.n)
			result := tc.tree.Root.String()
			if tc.expected != result {
				t.Errorf("Expected %s but received %s", tc.expected, result)
			}
		})
	}
}
