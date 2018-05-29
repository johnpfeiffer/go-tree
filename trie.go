package gotree

// TrieNode is a node in a prefix tree https://en.wikipedia.org/wiki/Trie
type TrieNode struct {
	Key      rune
	Children map[rune]*TrieNode
}

// TODO: is a string more readable and just as efficient as a rune?

// TrieInsert adds the string into the tree (adding nodes as necessary)
func TrieInsert(root *TrieNode, s string) {
	current := root
	for index := 0; index < len(s); index++ {
		if current.Children == nil {
			current.Children = make(map[rune]*TrieNode)
		}
		next, ok := current.Children[rune(s[index])]
		if !ok {
			next = &TrieNode{Key: rune(s[index])}
			current.Children[rune(s[index])] = next
		}
		current = next
	}
}

// Find returns if a string is in the Trie
func (t *TrieNode) Find(s string) bool {
	current := t
	index := 0
	for ; index < len(s); index++ {
		next, ok := current.Children[rune(s[index])]
		if !ok {
			break
		}
		current = next
	}
	if index == len(s) {
		return true
	}
	return false
}

// getWords gets all of the terminating strings in the Trie
func (t *TrieNode) getWords() []string {
	// leaf node is the base case
	if len(t.Children) == 0 {
		// do not return the root key that is the rune zero value
		if t.Key != 0 {
			return []string{string(t.Key)}
		}
		return []string{}
	}

	// TODO: refactor to enumerate alphabetically
	result := []string{}
	for _, r := range t.Children {
		result = append(result, r.getWords()...)
	}

	// do not return the root key that is the rune zero value
	if t.Key == 0 {
		return result
	}

	final := []string{}
	for _, s := range result {
		final = append(final, string(t.Key)+s)
	}
	return final
}

// TODO: remove unused?
// nextNode discovers if a new node is necessary in the tree or returns a pointer to the next node
func nextNode(children []*TrieNode, target rune) *TrieNode {
	// TODO: this is slow but allows enumeration alphabetically
	for _, c := range children {
		if target == c.Key {
			return c
		}
	}
	return nil
}
