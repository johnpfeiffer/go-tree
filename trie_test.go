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
	assertBoolean(t, false, root.Children['h'].EndOfWord)
	assertTrieKey(t, root.Children['h'].Children['i'].Key, 'i')
	assertBoolean(t, true, root.Children['h'].Children['i'].EndOfWord)

	assertBoolean(t, true, root.Find("hi"))
	assertBoolean(t, false, root.Find("high"))

	contents := root.getWords()
	assertNumber(t, "number of leaf nodes in the Trie", 1, len(contents))
	assertString(t, "Trie contents", "hi", contents[0])

	TrieInsert(root, "hello")

	assertBoolean(t, true, root.Find("hi"))
	assertBoolean(t, true, root.Find("hello"))
	shouldNotFind(t, root, []string{"", "he", "hell", "high", "hella"})
	contents = root.getWords()
	assertNumber(t, "number of leaf nodes in the Trie", 2, len(contents))
	assertString(t, "Trie contents", "hi", contents[0])
	assertString(t, "Trie contents", "hello", contents[1])
}

func TestTrieFind(t *testing.T) {
	root := &TrieNode{}                     // root has no key associated with it, only children
	assertBoolean(t, false, root.Find(""))  // empty trie finds nothing
	assertBoolean(t, false, root.Find("a")) // empty trie finds nothing

	TrieInsert(root, "a")
	assertBoolean(t, true, root.Find("a"))
	shouldNotFind(t, root, []string{"", "aa", "i", "am"}) // edge case, repetition, does not exist, too long

	TrieInsert(root, "and")
	assertBoolean(t, true, root.Find("and"))
	assertBoolean(t, false, root.Find("an"))              // will not find without an end-of-word
	assertBoolean(t, true, root.Find("a"))                // regression check
	shouldNotFind(t, root, []string{"", "aa", "i", "am"}) // edge case, repetition, does not exist, too long
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

func shouldNotFind(t *testing.T, root *TrieNode, words []string) {
	for _, word := range words {
		assertBoolean(t, false, root.Find(word))
	}
}

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
