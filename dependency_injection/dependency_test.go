package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGreet(t *testing.T) {
	var buffer bytes.Buffer
	Greet(os.Stdout, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
