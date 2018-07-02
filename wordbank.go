package gotree

import (
	"fmt"
)

// WordBank stores words for suggesting auto-completion
type WordBank interface {
	Insert(word string)
	Remove(word string)
	GetSuggestions(prefix string) []string
}

// MapWordBank is an implementation of a WordBank using a map
type MapWordBank struct {
	suggestions map[string][]string
}

// Insert stores a word in the wordbank so it can be retrieved for suggestions
func (m *MapWordBank) Insert(word string) {
	if m.suggestions == nil {
		m.suggestions = make(map[string][]string)
	}
	for i := len(word); i > 0; i-- {
		prefix := string(word[0:i])
		current, ok := m.suggestions[prefix]
		if !ok {
			m.suggestions[prefix] = []string{word}
		} else {
			m.suggestions[prefix] = append(current, word)
		}
	}
}

// Remove a word in the wordbank (it will no longer appear in suggestions)
func (m *MapWordBank) Remove(word string) {
	if m.suggestions == nil {
		return
	}
	for i := len(word); i > 0; i-- {
		prefix := string(word[0:i])
		current, ok := m.suggestions[prefix]
		if !ok {
			fmt.Println("WARNING: the word was not found:", word, "with prefix:", prefix)
			break
		} else {
			reduced, err := removeFromSlice(current, word)
			if err != nil {
				fmt.Println("WARNING: the word was not found:", word, "with prefix:", prefix)
			}
			m.suggestions[prefix] = reduced
		}
	}
}

// removeFromSlice returns an error if the target is not in the slice (so not idempotent)
func removeFromSlice(a []string, target string) ([]string, error) {
	found := false
	reduced := []string{}
	for i := 0; i < len(a); i++ {
		if a[i] == target {
			found = true
		} else {
			reduced = append(reduced, a[i])
		}
	}
	if !found {
		return nil, fmt.Errorf("in the slice could not find: %s", target)
	}
	return reduced, nil
}

// GetSuggestions returns all words that start with the prefix, for convenience returning empty rather than nil
func (m *MapWordBank) GetSuggestions(prefix string) []string {
	words, ok := m.suggestions[prefix]
	if !ok {
		return []string{}
	}
	return words
}
