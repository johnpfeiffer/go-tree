package gotree

import "testing"

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
