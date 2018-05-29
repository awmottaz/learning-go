package echoes

import "testing"

func BenchmarkEcho1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Echo1("a", "b", "c", "d", "e", "f", "g")
	}
}

func BenchmarkEcho2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Echo2("a", "b", "c", "d", "e", "f", "g")
	}
}

func BenchmarkEcho3(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Echo3("a", "b", "c", "d", "e", "f", "g")
	}
}
