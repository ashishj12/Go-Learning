package main

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "http://google.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
	}

	want := map[string]bool{
		"http://google.com":          true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}