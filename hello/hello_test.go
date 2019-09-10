package main

import "testing"

func TestHello(t *testing.T) {
	real := Hello()
	expected := "Hello, World"

	if real != expected {
		t.Errorf("real : '%s' expected '%s'", real, expected)
	}
}
