package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
	// In binary, dividing an integer by two is the same as dropping off the
	// last bit. So the population count of i is the same as i/2 if i is
	// even, and one greater if i is odd.
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Loop returns the population count (number of set bits) of x, using a
// loop.
func Loop(x uint64) int {
	var c uint64
	for i := 0; i < 64; i++ {
		c += (x >> uint(i)) & 1
	}
	return int(c)
}
