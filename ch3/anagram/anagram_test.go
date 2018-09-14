package anagram

import "testing"

func TestAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"not anagrams", args{"foo", "bar"}, false},
		{"simple anagram", args{"foo", "ofo"}, true},
		{"subset smaller", args{"foo", "fooo"}, false},
		{"subset larger", args{"fooo", "foo"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Anagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Anagram(%q, %q) = %v, want %v", tt.args.s1, tt.args.s2, got, tt.want)
			}
		})
	}
}
