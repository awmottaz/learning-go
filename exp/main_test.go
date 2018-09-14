package exp

import "testing"

func TestExp(t *testing.T) {
	type args struct {
		a int
		p int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{3, 2}, 9.},
		{"2", args{2, 4}, 16.},
		{"3", args{3, 7}, 2187.},
		{"4", args{-1, 2}, 1.},
		{"5", args{5, -3}, 0.008},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exp(tt.args.a, tt.args.p); got != tt.want {
				t.Errorf("exp(%v, %v) = %v, want %v", tt.args.a, tt.args.p, got, tt.want)
			}
		})
	}
}

func Test_expSlow(t *testing.T) {
	type args struct {
		a int
		p int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{3, 2}, 9.},
		{"2", args{2, 4}, 16.},
		{"3", args{3, 7}, 2187.},
		{"4", args{-1, 2}, 1.},
		{"5", args{5, -3}, 0.008},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expSlow(tt.args.a, tt.args.p); got != tt.want {
				t.Errorf("exp(%v, %v) = %v, want %v", tt.args.a, tt.args.p, got, tt.want)
			}
		})
	}
}

func BenchmarkExp(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Exp(17, 5000)
	}
}

func Benchmark_expSlow(t *testing.B) {
	for i := 0; i < t.N; i++ {
		expSlow(17, 5000)
	}
}
