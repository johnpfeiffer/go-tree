package gotree

import (
	"testing"
)

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

func assertNumber(t *testing.T, hint string, expected, result int) {
	t.Helper()
	if expected != result {
		t.Error("\nExpected", hint, ":", expected, "\nReceived:", result)
	}
}

func assertNil(t *testing.T, thing interface{}) {
	t.Helper()
	if thing != nil {
		t.Error("should be nil, instead it was", thing)
	}
}

func assertStringSlices(t *testing.T, hint string, expected, actual []string) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Error("Expected length:", len(expected), "but received length:", len(actual))
		return
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error("Expected value at index:", i, "was:", expected[i], "but received:", actual[i])
		}
	}
}
