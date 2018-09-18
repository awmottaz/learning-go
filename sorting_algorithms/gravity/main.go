package main

import (
	"fmt"
)

/*
This sorting algorithm works by imagining the following scenario: you have an
abacus with as many horizontal dowels as the largest number in the list you wish
to sort. For each number in your list, create a column of beads so that the beads
are added to the top-most dowel first, then the next dowel down, and so on until
the number of beads in the column equals the number in your list.

For example, if my list was {2, 1, 3}, my abacus setup would look like this:

	|--*--*--*--|
	|--*-----*--|
	|--------*--|

To sort the numbers, rotate the abacus 90 degrees so that the beads fall to one
side, then read from top to bottom.

	---------
	 |  |  |
	 *  |  |
	 |  |  |
	 *  *  |
	 |  |  |
	 *  *  *
	 |  |  |
	---------
*/

func main() {
	data := []int{3, 1, 12, 42, 8, 96, 3, 14, 15, 7, 28, 90}
	var abacus []int
	var sorted []int

	// Add beads to the abacus for each value in the list we're sorting...
	for _, v := range data {
		// ...starting with the top-most dowel.
		for i := 0; i < v; i++ {
			// The index i corresponds to the dowels of the abacus
			// from top to bottom. We need as many dowels as the
			// largest number to be sorted. If there are not enough
			// dowels, add one.
			if len(abacus) <= i {
				abacus = append(abacus, 0)
			}

			// Add a bead to this dowel.
			abacus[i]++
		}
	}

	// To read off the sorted values (after rotating the abacus), count the
	// number of beads in the column and then remove them. The first dowel has
	// len(data) beads, so this process terminates when the first dowel runs
	// out of beads.
	for abacus[0] > 0 {
		val := countCol(abacus)
		sorted = append(sorted, val)
		abacus = rmvCol(abacus)
	}

	fmt.Printf("in\n\t%v\nout\n\t%v\n", data, sorted)
}

// countCol counts the number of beads in the left-most column of the given abacus.
func countCol(a []int) int {
	var c int
	for _, v := range a {
		if v > 0 {
			c++
		}
	}
	return c
}

// rmvCol returns a new abacus with the beads from the left-most column removed.
func rmvCol(a []int) []int {
	for i, v := range a {
		if v <= 0 {
			a[i] = 0
		} else {
			a[i] = v - 1
		}
	}
	return a
}
