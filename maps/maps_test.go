package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, got := dictionary.Search("else")

		if got == nil {
			t.Fatal("Expected to get an error")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "This is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is just a test"}
		word := "test"
		definition := "This is a new definition"

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, "This is just a test")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "This is just a test"

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

	t.Run("Existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is just a test"}
		word := "test"
		definition := "This is a new test definition"

		err := dictionary.Update(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is just a test"}
		word := "test"

		err := dictionary.Delete(word)

		assertError(t, err, nil)

		_, err = dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})
	t.Run("Non-existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"

		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
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
		t.Errorf("Got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("Should find added word:", err)
	}
	assertStrings(t, got, definition)
}
