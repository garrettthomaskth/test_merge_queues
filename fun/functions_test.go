package fun

import "testing"

func TestConcatinateStringsWithSpace(t *testing.T) {
	testOutput := ConcatinateStringsWithSpace("1", "23", "4")
	expectedOutput := "1 23 4"
	if testOutput != expectedOutput {
		t.Fatalf("expected %s, got %s", expectedOutput, testOutput)
	}
}