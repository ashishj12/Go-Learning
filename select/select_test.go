package main

// import (
// 	"testing"
// 	"time"
// )

// func TestRacer(t *testing.T) {
// 	slowURL := "http://www.facebook.com"
// 	fastURL := "http://www.github.com"

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }

// func TestRacer(t *testing.T) {

// 	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(20 * time.Millisecond)
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	slowURL := slowServer.URL
// 	fastURL := fastServer.URL

// 	want := fastURL
// 	got := Racer(slowURL, fastURL)

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}

// 	slowServer.Close()
// 	fastServer.Close()
// }

// func TestRac(t *testing.T) {
// 	t.Run("compares speed of servers, returning the url of the fastest one", func(t *testing.T) {
// 		slowServer := makeDelayedServer(20 * time.Millisecond)
// 		fastServer := makeDelayedServer(0 * time.Millisecond)

// 		defer slowServer.Close()
// 		defer fastServer.Close()

// 		slowURL := slowServer.URL
// 		fastURL := fastServer.URL

// 		want := fastURL
// 		got, _ := Racer(slowURL, fastURL)

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})

// 	t.Run("returns an error if a server doesn't responed within 10s", func(t *testing.T) {
// 		serverA := makeDelayedServer(11 * time.Second)
// 		serverB := makeDelayedServer(12 * time.Second)

// 		defer serverA.Close()
// 		defer serverB.Close()

// 		_, err := Racer(serverA.URL, serverB.URL)

// 		if err == nil {
// 			t.Error("expected an error but didn't get one")
// 		}

// 	})

// }
