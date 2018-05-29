// Echo5 prints the index and value of each command line argument, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[1:] {
		fmt.Printf("Argument %v is: %v\n", i, v)
	}
}
