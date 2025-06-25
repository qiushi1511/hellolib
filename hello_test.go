package hellolib

import "testing"

func TestGreet(t *testing.T) {
	result := Greet("World")
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Greet() = %q, want %q", result, expected)
	}
}

func TestWarmGreet(t *testing.T) {
	result := WarmGreet("World")
	expected := "Hello, World! Welcome to the world of Go!"
	if result != expected {
		t.Errorf("WarmGreet() = %q, want %q", result, expected)
	}
}
