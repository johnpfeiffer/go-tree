package gotree

// WordBank stores words for suggesting auto-completion
type WordBank interface {
	Insert(word string)
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
	for i := 0; i < len(word); i++ {
		prefix := string(word[:i])
		current, ok := m.suggestions[prefix]
		if !ok {
			m.suggestions[prefix] = []string{word}
		} else {
			m.suggestions[prefix] = append(current, word)
		}
	}
}

// GetSuggestions returns all words that start with the prefix, for convenience returning empty rather than nil
func (m *MapWordBank) GetSuggestions(prefix string) []string {
	words, ok := m.suggestions[prefix]
	if !ok {
		return []string{}
	}
	return words
}
