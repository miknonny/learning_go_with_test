package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(time.Second * 3)
	}
	fmt.Fprintln(w, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
