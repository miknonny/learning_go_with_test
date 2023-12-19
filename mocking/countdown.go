package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(10 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}

	fmt.Fprintln(w, finalWord)
}

func main() {
	var sleep DefaultSleeper
	Countdown(os.Stdout, &sleep)
}
