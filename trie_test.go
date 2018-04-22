package gotree

import (
	"fmt"
	"testing"
)

func TestNextNode(t *testing.T) {
	var testCases = []struct {
		children []*TrieNode
		target   rune
		expected rune
	}{
		{children: nil, target: 'a', expected: 0},
		{children: []*TrieNode{}, target: 'a', expected: 0},
		{children: []*TrieNode{&TrieNode{key: 'a'}}, target: 'a', expected: 'a'},
		{children: []*TrieNode{&TrieNode{key: 'a'}, &TrieNode{key: 'b'}}, target: 'a', expected: 'a'},
		{children: []*TrieNode{&TrieNode{key: 'a'}, &TrieNode{key: 'b'}}, target: 'b', expected: 'b'},
		{children: []*TrieNode{&TrieNode{key: 'a'}, &TrieNode{key: 'b'}}, target: 'c', expected: 0},
		{children: generateTrieChildren(t, []rune{'a', 'b', 'c', 'd', 'e', 'A', 'B', 'Z'}), target: 'A', expected: 'A'},
		{children: generateTrieChildren(t, []rune{'a', 'b', 'c', 'd', 'e', 'A', 'B', 'Z'}), target: 'z', expected: 0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v finding %v", tc.children, string(tc.target)), func(t *testing.T) {
			var result rune
			n := nextNode(tc.children, tc.target)
			if n != nil {
				result = n.key
			}
			if tc.expected != result {
				t.Errorf("Expected %v but received %v", tc.expected, result)
			}
		})
	}
}

// helpers
func generateTrieChildren(t *testing.T, a []rune) []*TrieNode {
	t.Helper()
	result := []*TrieNode{}
	for _, r := range a {
		result = append(result, &TrieNode{key: r})
	}
	return result
}
