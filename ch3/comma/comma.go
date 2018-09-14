package comma

import (
	"bytes"
	"strings"
)

// Comma inserts commas into number strings.
func Comma(s string) string {
	var (
		buf bytes.Buffer
		dec string
	)

	decInd := strings.Index(s, ".")
	if decInd >= 0 {
		dec = s[decInd:]
		s = s[:decInd]
	}

	for i, r := range s {
		if i > 0 && len(s[i:])%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(r)
	}

	buf.WriteString(dec)
	return buf.String()
}
