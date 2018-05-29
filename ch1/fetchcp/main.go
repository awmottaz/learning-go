// Fetchcp performs a GET request of the given URL and
// copies the contents to os.Stdout
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing URL")
		os.Exit(1)
	}
	//!+exercise 1.8
	url := os.Args[1]
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	//!-exercise 1.8
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	//!+exercise 1.7
	//!+exercise 1.9
	fmt.Printf("GET - %v\n", resp.Status)
	//!-exercise 1.9
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error copying to std out: %v\n", err)
		os.Exit(1)
	}
	//!-exercise 1.7
}
