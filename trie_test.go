package gotree

import (
	"fmt"
	"testing"
)

func TestTrieInsertBasic(t *testing.T) {
	root := &TrieNode{} // root has no key associated with it, only children
	TrieInsert(root, "hi")
	assertTrieKey(t, root.Key, 0)

	assertTrieKey(t, root.Children['h'].Key, 'h')
	assertTrieKey(t, root.Children['h'].Children['i'].Key, 'i')

	assertBoolean(t, true, root.Find("hi"))
	assertBoolean(t, false, root.Find("high"))

	contents := root.getWords()
	assertNumber(t, "number of leaf nodes in the Trie", 1, len(contents))
	assertString(t, "Trie contents", "hi", contents[0])

	TrieInsert(root, "hello")

	assertBoolean(t, true, root.Find("hi"))
	assertBoolean(t, true, root.Find("hello"))
	assertBoolean(t, true, root.Find("he"))
	assertBoolean(t, true, root.Find("hell"))
	assertBoolean(t, false, root.Find("high"))
	assertBoolean(t, false, root.Find("hella"))
	contents = root.getWords()
	assertNumber(t, "number of leaf nodes in the Trie", 2, len(contents))

	assertString(t, "Trie contents", "hi", contents[0])
	assertString(t, "Trie contents", "hello", contents[1])
}
func TestTrieNextNode(t *testing.T) {
	var testCases = []struct {
		children []*TrieNode
		target   rune
		expected rune
	}{
		{children: nil, target: 'a', expected: 0},
		{children: []*TrieNode{}, target: 'a', expected: 0},
		{children: []*TrieNode{&TrieNode{Key: 'a'}}, target: 'a', expected: 'a'},
		{children: []*TrieNode{&TrieNode{Key: 'a'}, &TrieNode{Key: 'b'}}, target: 'a', expected: 'a'},
		{children: []*TrieNode{&TrieNode{Key: 'a'}, &TrieNode{Key: 'b'}}, target: 'b', expected: 'b'},
		{children: []*TrieNode{&TrieNode{Key: 'a'}, &TrieNode{Key: 'b'}}, target: 'c', expected: 0},
		{children: generateTrieChildren(t, []rune{'a', 'b', 'c', 'd', 'e', 'A', 'B', 'Z'}), target: 'A', expected: 'A'},
		{children: generateTrieChildren(t, []rune{'a', 'b', 'c', 'd', 'e', 'A', 'B', 'Z'}), target: 'z', expected: 0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v finding %v", tc.children, string(tc.target)), func(t *testing.T) {
			var result rune
			n := nextNode(tc.children, tc.target)
			if n != nil {
				result = n.Key
			}
			assertTrieKey(t, tc.expected, result)
		})
	}
}

// helpers

func generateTrieChildren(t *testing.T, a []rune) []*TrieNode {
	t.Helper()
	result := []*TrieNode{}
	for _, r := range a {
		result = append(result, &TrieNode{Key: r})
	}
	return result
}

func assertTrieKey(t *testing.T, expected, actual rune) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %#v but received %#v", expected, actual)
	}
}

func assertBoolean(t *testing.T, expected, actual bool) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %v but received %v", expected, actual)
	}
}

func assertString(t *testing.T, hint, expected, actual string) {
	t.Helper()
	if expected != actual {
		t.Errorf("%s Expected %#v but received %#v", hint, expected, actual)
	}
}
