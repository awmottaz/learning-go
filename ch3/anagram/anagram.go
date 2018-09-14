package anagram

import (
	"strings"
)

// Anagram determines whether s1 and s2 are anagrams of each other.
func Anagram(s1, s2 string) bool {
	for _, r := range s1 {
		i := strings.IndexRune(s2, r)
		if i < 0 {
			return false
		}
		s2 = strings.Replace(s2, string(r), "", 1)
	}
	return len(s2) == 0
}
