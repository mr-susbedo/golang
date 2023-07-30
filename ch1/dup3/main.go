// entire input into memory in one big gulp, split it into lines all at once, then process the lines
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _,line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
		
	for text, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n",count,text)
		}
	}
}