package main

import "testing"

func TestHello(t *testing.T) {
	// What we got
	got := Hello("GDSC")
	want := "Hello, GDSC"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
