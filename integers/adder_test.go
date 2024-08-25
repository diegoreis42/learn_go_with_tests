package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	assertResult(t, sum, expected)
}

func TestSub(t *testing.T) {
	result := Sub(4, 2)
	expected := 2

	assertResult(t, result, expected)
}

func assertResult(t testing.TB, got, expected int) {
	t.Helper()

	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
}
