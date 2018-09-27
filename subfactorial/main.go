// subfactorial computes the number of derangements of n objects. It requires
// one input, a non-negative integer, and returns the subfactorial of that number.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing input")
	}
	in := os.Args[1]
	n, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	if n < 0 {
		log.Fatalf("input must be a non-negative integer, got %v", in)
	}
	fmt.Println(subfactorial(n))
}

func subfactorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}
	return (n - 1) * (subfactorial(n-1) + subfactorial(n-2))
}
