package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop(uint64(i))
	}
}

func TestLoop(t *testing.T) {
	for i := uint64(0); i < 20000; i++ {
		if Loop(i) != PopCount(i) {
			t.Fail()
		}
	}
}

func BenchmarkArgshift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Argshift(uint64(i))
	}
}

func TestArgshift(t *testing.T) {
	for i := uint64(0); i < 20000; i++ {
		if Argshift(i) != PopCount(i) {
			t.Fail()
		}
	}
}

func BenchmarkClearbits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Clearbits(uint64(i))
	}
}

func TestClearbits(t *testing.T) {
	for i := uint64(0); i < 20000; i++ {
		if Clearbits(i) != PopCount(i) {
			t.Fail()
		}
	}
}
