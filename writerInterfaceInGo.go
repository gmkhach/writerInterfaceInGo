package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// These two accomplish the same thing. In fact, Println is actually calling the same exact line that follows after it.
	fmt.Println("Hello, world")
	fmt.Fprintln(os.Stdout, "Hello, world")

	// Here is another function that accomplishes the same exact thing by using a Writer type os.Stdout:
	io.WriteString(os.Stdout, "Hello, world")
}