package main

import (
	"fmt"
)

func main() {
	data := []int{3, 1, 12, 42, 8, 96, 3, 14, 15, 7, 28, 90}
	var abacus []int
	var sorted []int

	for _, v := range data {
		for i := 0; i < v; i++ {
			// abacus[i]++
			if len(abacus) <= i {
				abacus = append(abacus, 1)
			} else {
				abacus[i]++
			}
		}
	}

	for abacus[0] > 0 {
		next := numNonzero(abacus)
		sorted = append(sorted, next)
	}

	fmt.Println(sorted)
}

func numNonzero(a []int) int {
	var sum int
	for i, v := range a {
		if v > 0 {
			sum++
			a[i]--
		}
	}
	return sum
}
