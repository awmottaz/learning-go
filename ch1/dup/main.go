// Dup prints the filename, count, and text of lines
// that appear more than once in the named input files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// filecounts maps the filename to a map of lines to their counts
	perFileCounts := make(map[string]map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
		}
		perFileCounts[filename] = make(map[string]int)
		for _, line := range strings.Split(string(data), "\n") {
			perFileCounts[filename][line]++
		}
	}
	for filename, fc := range perFileCounts {
		for line, count := range fc {
			if count > 1 {
				fmt.Printf("In file %v\n\t%d\t%s\n", filename, count, line)
			}
		}
	}
}
