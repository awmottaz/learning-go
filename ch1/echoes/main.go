// Package echoes provides multiple echo implementations.
package echoes

import (
	"fmt"
	"strings"
)

// Echo1 is a copy of the echo1 listing from D&K.
// See https://github.com/adonovan/gopl.io/tree/master/ch1/echo1.
func Echo1(a ...string) string {
	var s, sep string
	for i := 1; i < len(a); i++ {
		s += sep + a[i]
		sep = " "
	}
	return fmt.Sprint(s)
}

// Echo2 is a copy of the echo2 listing from D&K.
// See https://github.com/adonovan/gopl.io/blob/master/ch1/echo2.
func Echo2(a ...string) string {
	s, sep := "", ""
	for _, arg := range a[1:] {
		s += sep + arg
		sep = " "
	}
	return fmt.Sprint(s)
}

// Echo3 is a copy of the echo3 listing from D&K.
// See https://github.com/adonovan/gopl.io/blob/master/ch1/echo3.
func Echo3(a ...string) string {
	return fmt.Sprint(strings.Join(a[1:], " "))
}
