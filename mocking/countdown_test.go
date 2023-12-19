package main

import (
	"bytes"
	"testing"
)

// Imagine having a code with time.sleep or a request that takes a long time to run.
// such piece of code is the ideal candidate for mocking so we dont spend valuable
// time waiting for the code to execute.
func TestCount(t *testing.T) {
	var buf bytes.Buffer

	Countdown(&buf)

	got := buf.String()
	want := `3
2
1
Go!
`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
