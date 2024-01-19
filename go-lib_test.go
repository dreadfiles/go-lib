package lib

import "testing"

func TestAdd(t *testing.T) {
	a := 1
	b := 2

	expected := a + b
	got := Add(a, b)

	if got != expected {
		t.Errorf("Add(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}

func TestSubstract(t *testing.T) {
	a := 1
	b := 2

	expected := a - b
	got := Substract(a, b)

	if got != expected {
		t.Errorf("Substract(%d, %d) = %d, didn't return %d", a, b, got, expected)
	}
}
