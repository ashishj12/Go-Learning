package iteration

import "testing"

//test case for a function that return 'a' character 5 times
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %s but got %s", repeated, expected)
	}
}
