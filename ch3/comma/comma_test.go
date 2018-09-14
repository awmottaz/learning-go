package comma

import "testing"

func TestComma(t *testing.T) {
	tests := [][2]string{
		{"123", "123"},
		{"1234", "1,234"},
		{"1234567891234", "1,234,567,891,234"},
		{"12345.6789", "12,345.6789"},
	}
	for _, test := range tests {
		out := Comma(test[0])
		if out != test[1] {
			t.Errorf("Got %v, expected %v\n", out, test[1])
		}
	}
}
