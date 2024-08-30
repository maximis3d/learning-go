package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Max", "")
		want := "Hello, Max"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, world', when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say hello in Spanish", func(t *testing.T) {
		got := Hello("Max", "Spanish")
		want := "Hola, Max"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say hello in French", func(t *testing.T) {
		got := Hello("Max", "French")
		want := "Bonjour, Max"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
