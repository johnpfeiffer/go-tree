package gotree

// TrieNode is a node in a prefix tree https://en.wikipedia.org/wiki/Trie
type TrieNode struct {
	key      rune
	children []*TrieNode
}

// TODO: is a string more readable and just as efficient as a rune?

/*
// TrieInsert adds the string into the tree (adding nodes as necessary)
func TrieInsert(root *TrieNode, s string) {
	current := root
	index := 0

	for {
		if index == len(s)-1 {
			break
		}
		next := nextNode(current.children, rune(s[index]))
		if next == nil {
			n := &TrieNode{key: rune(s[index])}
			current.children = append(current.children, n)
		}
		current = next
	}
}

*/

// nextNode discovers if a new node is necessary in the tree or returns a pointer to the next node
func nextNode(children []*TrieNode, target rune) *TrieNode {
	for _, c := range children {
		if target == c.key {
			return c
		}
	}
	return nil
}
