package hellolib

import "testing"

func TestGreet(t *testing.T) {
	result := Greet("World")
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Greet() = %q, want %q", result, expected)
	}
}
