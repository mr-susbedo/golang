package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " ")) // Modify the echo program to also print os.Args[0], the name of the command that invoked it.

	// Modify the echo program to print the index and value of each of its arguments, one per line.
	for idx, arg := range os.Args {
		str := fmt.Sprintf("Index: %d, value: %s", idx, arg)
		fmt.Println(str)
	}
}
