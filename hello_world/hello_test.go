package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Max")
	want := "Hello, Max"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
