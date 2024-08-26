package main

import "testing"

func TestSearch(t *testing.T) {
	dictionay := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dictionay.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unkown word", func(t *testing.T) {
		_, err := dictionay.Search("unkown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionay := Dictionary{}
		word := "test"
		body := "this is just a test"

		dictionay.Add(word, body)

		assertDefinition(t, dictionay, word, body)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		body := "this is just a test"
		dictionay := Dictionary{word: body}
		err := dictionay.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionay, word, body)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionay Dictionary, key, definition string) {
	t.Helper()

	got, err := dictionay.Search(key)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
