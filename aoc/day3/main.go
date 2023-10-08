package main

import (
	"fmt"
	"unicode"
)

func getPriorityValue(char rune) int {
	if unicode.IsLower(char) {
		return int(char-'a') + 1
	} else if unicode.IsUpper(char) {
		return int(char-'A') + 27
	}
	return 0
}

func main() {
	chars := []rune{'a', 'A', 'z', 'Z', 'f', 'F'}

	for _, char := range chars {
		fmt.Printf("Alphabet: %c, Priority: %d\n", char, getPriorityValue(char))
	}
}
