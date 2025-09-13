package main

import (
	// "fmt"
	"testing"
)

//pass test case
// func TestHello(t *testing.T) {

// 	got := Hello()
// 	want := "Hello, world"
// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

//fail test case
// func TestHe(t *testing.T) {
// 	got := Hello()
// 	want := "Hello World"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// func TestHello(t *testing.T) {

// 	got := Hello("golang")
// 	want := "Hello, go"

// 	if got != want {
// 		fmt.Printf("got %q want %q\n", got, want)
// 	}

// }

func TestHello(t *testing.T) {
	t.Run("Saying hello to golang", func(t *testing.T) {
		got := Hello("User")
		want := "Hello, User"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
