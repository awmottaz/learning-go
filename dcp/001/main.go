// Daily Coding Problem 001
//
// Given a list of numbers and a number k, return whether any two numbers
// from the list add up to k.
//
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
//
// Bonus: Can you do this in one pass?

package main

// My first solution takes a naive approach to solving the problem by simply
// checking every pair of numbers from the list to see if they hit the target
// number. This solution takes O(n^2) time.
func solutionOne(list []int, target int) bool {
	for i, x := range list {
		for _, y := range list[i+1:] {
			if x+y == target {
				return true
			}
		}
	}
	return false
}

// The second approach assumes that the given input list is already sorted,
// but is able to solve the problem with a single pass through the array.
// It works by picking out the first and last numbers from the list and shifting
// the selected pair towards the center depending on whether the sum is above or
// below the target. For example, given
//
// 	list: [3, 7, 10, 15], target: 17
//
// we begin with the pair {3, 15}. Since their sum is larger than 17, we exchange
// the larger number, 15, with the next largest, 10. Now the sum of 3 and 10 is lower
// than the target, so we exchange the 3 for the next smallest number, 7, and the
// pair {7, 10} hits the target.
func solutionTwo(list []int, target int) bool {
	left, right := 0, len(list)-1
	for left < right {
		sm, bg := list[left], list[right]
		if sm+bg == target {
			return true
		} else if sm+bg > target {
			right--
		} else {
			left++
		}
	}
	return false
}

// The final solution is an iteration on the previous solution. In principle it
// works the same way, but the ordered list upon which it operates is sort of built
// up as we go.
func finalSolution(list []int, target int) bool {
	if len(list) < 2 {
		return false
	}
	if len(list) == 2 {
		return list[0]+list[1] == target
	}
	a, b := list[0], list[1]
	if a+b == target {
		return true
	}
	// Ensure a <= b
	if b < a {
		a, b = b, a
	}
	for i := 2; i < len(list); i++ {
		c := list[i]
		// Ensure a <= b <= c.
		if c < a {
			a, b, c = c, a, b
		} else if c < b {
			b, c = c, b
		}
		// Take the next step towards a pair that will work.
		if a+c == target {
			return true
		} else if a+c < target {
			a, b = b, c
		} // else a, b = a, b
	}
	return false
}
