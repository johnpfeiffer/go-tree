package gotree

import (
	"fmt"
	"testing"
)

func TestMapWordBankInsertBasic(t *testing.T) {
	var w MapWordBank
	assertNumber(t, "word suggestion empty", 0, len(w.GetSuggestions("h")))
	assertNumber(t, "word map empty", 0, len(w.suggestions))

	w.Insert("hi")
	actualFirst := w.GetSuggestions("h")
	assertNumber(t, "word inserted suggestion count", 1, len(actualFirst))
	assertString(t, "word inserted", "hi", actualFirst[0])

	assertNumber(t, "map count", 2, len(w.suggestions))
	assertString(t, "map item", "hi", w.suggestions["h"][0])
	assertString(t, "map item", "hi", w.suggestions["hi"][0])
}

func TestMapWordBankInsertAdvanced(t *testing.T) {
	var w MapWordBank
	w.Insert("hi")
	assertNumber(t, "word inserted suggestion count", 1, len(w.GetSuggestions("h")))
	assertString(t, "word inserted", "hi", w.GetSuggestions("h")[0])

	w.Insert("highest")
	actual := w.GetSuggestions("h")
	assertNumber(t, "word inserted suggestion count", 2, len(actual))
	assertString(t, "word inserted", "hi", actual[0])
	assertString(t, "word inserted", "highest", actual[1])

	assertNumber(t, "map count", 7, len(w.suggestions))
}

func TestRemoveFromSliceSuccess(t *testing.T) {
	var testCases = []struct {
		a        []string
		target   string
		expected []string
	}{
		{a: []string{"hi"}, target: "hi", expected: []string{}},
		{a: []string{"high", "hi"}, target: "hi", expected: []string{"high"}},
		{a: []string{"hi", "high"}, target: "hi", expected: []string{"high"}},
		{a: []string{"", "hi", "high"}, target: "hi", expected: []string{"", "high"}},
		{a: []string{"", "hi", "high"}, target: "", expected: []string{"hi", "high"}},
		// TODO: UTF8
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v removes %v", tc.a, tc.target), func(t *testing.T) {
			result, err := removeFromSlice(tc.a, tc.target)
			assertNil(t, err)
			assertStringSlices(t, "remove from slice", tc.expected, result)
		})
	}

}

func TestRemoveFromSliceErrors(t *testing.T) {
	var testCases = []struct {
		a        []string
		target   string
		expected string
	}{
		{a: nil, target: "hi", expected: "in the slice could not find: hi"},
		{a: []string{}, target: "hi", expected: "in the slice could not find: hi"},
		{a: []string{"he"}, target: "hi", expected: "in the slice could not find: hi"},
		{a: []string{"h", "high", "hit"}, target: "hi", expected: "in the slice could not find: hi"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v removes %v", tc.a, tc.target), func(t *testing.T) {
			result, err := removeFromSlice(tc.a, tc.target)
			assertNumber(t, "remove from slice errorcases", 0, len(result))
			assertString(t, "remove from slice errorcases", tc.expected, err.Error())
		})
	}
}
