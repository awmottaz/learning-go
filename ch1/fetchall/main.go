// Fetchall fetches URLs in parallel and reports their times and sizes.
//
// Exercise 1.10
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dustin/go-humanize"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, urlstr := range os.Args[1:] {
		go fetch(urlstr, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(urlstr string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(urlstr)
	if err != nil {
		secs := time.Since(start).Seconds()
		ch <- fmt.Sprintf("%.2fs  %v", secs, err) // send to channel ch
		return
	}

	u, err := url.Parse(urlstr)
	if err != nil {
		ch <- fmt.Sprintf("while parsing url %s: %v", urlstr, err)
		return
	}
	outfilename := u.Hostname() + ".html"
	outfile, err := os.Create(outfilename)
	if err != nil {
		ch <- fmt.Sprintf("while creating file %s.html: %v", urlstr, err)
		return
	}
	nbytes, err := io.Copy(outfile, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", urlstr, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %6s  %s", secs, humanize.Bytes(uint64(nbytes)), urlstr)
}
