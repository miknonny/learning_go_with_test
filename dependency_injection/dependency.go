package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(w io.Writer, name string) {
	fmt.Fprint(w, name)
}

func main() {
	Greet(os.Stdout, "This is johnny")
}
