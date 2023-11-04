// Print the content found at URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mr-susbedo/golang/utils"
)

func main() {
	for _, url := range os.Args[1:] {

		url = utils.EnsurePrefix(url, utils.HTTP)

		resp, err := http.Get(url)

		if err != nil {
			if resp != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v, status code: %d\n", err, resp.StatusCode)
			} else {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			}
			os.Exit(1)
		}

		fmt.Printf("Status code: %d\n", resp.StatusCode)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}

	}
}
