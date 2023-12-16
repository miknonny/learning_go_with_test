package dictionary

import (
	"testing"
)

// A map value is a pointer to a runtime.hmap structure.
func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	// default test.
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	// Triggering the error case.
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown word")

		assertError(t, got, ErrNotFound)

	})

	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just another test"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existiing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("update definition", func(t *testing.T) {
		word := "test"
		definition := "random definition of a word."
		dict := Dictionary{word: definition}

		newDefinition := "another definition of the word."
		err := dict.Update(word, newDefinition)
		assertError(t, err, nil)

		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "testing this word"

		dict := Dictionary{}

		err := dict.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)

	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "my definition"

	d := Dictionary{word: definition}

	d.Delete(word)

	_, err := d.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal(ErrNotFound)
	}

	assertStrings(t, got, definition)

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
