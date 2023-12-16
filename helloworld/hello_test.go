package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("if empty name ", func(t *testing.T) {
		got := Hello("", English)
		want := "Hello, chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("if empty name is supplied", func(t *testing.T) {
		got := Hello("", English)
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", Spanish)
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", French)
		want := "Bonjour, Elodie"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Write the test.
// make the compiler pass.
// run the test.
// write enough code to make the test pass.
// refactor code.
