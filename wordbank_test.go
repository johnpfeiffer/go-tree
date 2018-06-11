package gotree

import (
	"testing"
)

func TestMapWordBankInsertBasic(t *testing.T) {
	var w MapWordBank
	w.GetSuggestions("h")
	assertNumber(t, "word inserted suggestion count empty", 0, len(w.GetSuggestions("h")))

	w.Insert("hi")
	actualFirst := w.GetSuggestions("h")
	assertNumber(t, "word inserted suggestion count", 1, len(actualFirst))
	assertString(t, "word inserted", "hi", actualFirst[0])
}

func TestMapWordBankInsertAdvanced(t *testing.T) {
	var w MapWordBank
	w.Insert("hi")
	assertNumber(t, "word inserted suggestion count", 1, len(w.GetSuggestions("h")))
	assertString(t, "word inserted", "hi", w.GetSuggestions("h")[0])

	w.Insert("hello")
	actual := w.GetSuggestions("h")
	assertNumber(t, "word inserted suggestion count", 2, len(actual))
	assertString(t, "word inserted", "hi", actual[0])
	assertString(t, "word inserted", "hello", actual[1])
}
