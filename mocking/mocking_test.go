package main

// import (
// 	"bytes"
// 	"testing"
// )

// // import (
// // 	"bytes"
// // 	"testing"
// // )

// // // func TestCountdown(t *testing.T) {
// // // 	buffer := &bytes.Buffer{}

// // // 	Countdown(buffer)

// // // 	got := buffer.String()
// // // 	want := "3"
// // // 	if got != want {
// // // 		t.Errorf("got %q want %q", got, want)
// // // 	}
// // // }

// // //update test case

// // func TestCountdown(t *testing.T) {
// // 	buffer := &bytes.Buffer{}

// // 	Countdown(buffer)

// // 	got := buffer.String()
// // 	want := `3
// // 2
// // 1
// // Go!`

// // 	if got != want {
// // 		t.Errorf("got %q want %q", got, want)
// // 	}
// // }


// // func TestCountdown(t *testing.T) {
// // 	buffer := &bytes.Buffer{}
// // 	spySleeper := &spySleeper{}

// // 	Countdown(buffer, spySleeper)

// // 	got := buffer.String()
// // 	want := `3
// // 2
// // 1
// // Go!`

// // 	if got != want {
// // 		t.Errorf("got %q want %q", got, want)
// // 	}

// // 	if spySleeper.Calls != 3 {
// // 		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
// // 	}
// // }