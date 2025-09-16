package main

import "testing"

// func TestMap(t *testing.T) {

// 	dictionary := map[string]string{"test": "this is just an simple test"}
// 	got := Search(dictionary, "test")
// 	want := "this is just an simple test"

// 	if got != want {
// 		t.Errorf("got %q want %q given, %q", got, want, "test")
// 	}

// }

// func TestSearch(t *testing.T) {
// 	dictionary := Dictionary{"test":"this is just an simple test"}
// 	got := dictionary.Search("test")
// 	want := "this is just an simple test"

// 	assertStrings(t, got, want)

// }

// func assertStrings(t testing.TB, got, want string) {
// 	t.Helper()

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}

// }

//test case where word in known or unknown at the same time

// func TestSearch(t *testing.T) {
// 	dictionary := Dictionary{"test": "this is just a test"}

// 	t.Run("known word", func(t *testing.T) {
// 		got, _ := dictionary.Search("test")
// 		want := "this is just a test"

// 		assertStrings(t, got, want)
// 	})

// 	t.Run("unknown word", func(t *testing.T) {
// 		_, err := dictionary.Search("unknown")
// 		want := "could not find the word you were looking for"

// 		if err == nil {
// 			t.Fatal("expected to get an error.")
// 		}

// 		assertStrings(t, err.Error(), want)
// 	})
// }

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is an simple test")

	want := "this is an simple test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, want)
}
