package main

import "testing"

func TestFoo(t *testing.T) {
	if got, want := foo(), "Hello, World!"; got != want {
		t.Errorf("foo() = %q; want %q", got, want)
	}
}
