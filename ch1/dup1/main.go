// program to FIND DUPLICATE LINES
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		text := input.Text()
		counts[text]++
	}

	fmt.Println("Duplicate Occurrences:")
	for text, count := range counts {
		if(count > 1) {
			fmt.Printf("%d\t%s\n",count,text)
		}
	}

}