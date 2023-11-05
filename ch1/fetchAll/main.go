// fetchAll fetches URLs in parallel and reports their times and sizes.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	for i := 0; i < 2; i++ {
		fetchAndReport()
		fmt.Println("Run completed: ", i+1)
	}

}

func fetchAndReport() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	// Print cache-related headers
	fmt.Println("URL:", url)
	fmt.Println("Cache-Control:", resp.Header.Get("Cache-Control"))
	fmt.Println("ETag:", resp.Header.Get("ETag"))
	fmt.Println("Last-Modified:", resp.Header.Get("Last-Modified"))
	fmt.Println()

	resp.Body.Close() // don't leak resources (O_O)

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)

}
