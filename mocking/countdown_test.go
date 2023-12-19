package main

import (
	"bytes"
	"testing"
)

// Imagine having a code with time.sleep or a request that takes a long time to run.
// such piece of code is the ideal candidate for mocking so we dont spend valuable
// time waiting for the code to execute.
// Always remember that slow feedback ruins developer productivity.

// The spy is  a kind of mock that helps us track how many times sleep is called
// tracks how dependencies are used.
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCount(t *testing.T) {
	var buf bytes.Buffer
	var spySleeper SpySleeper

	Countdown(&buf, &spySleeper)

	got := buf.String()
	want := `3
2
1
Go!
`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not  enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}
