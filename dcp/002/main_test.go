package dcp002

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name:       "1",
			args:       args{[]int{1, 2, 3, 4, 5}},
			wantResult: []int{120, 60, 40, 30, 24},
		},
		{
			name:       "2",
			args:       args{[]int{3, 2, 1}},
			wantResult: []int{2, 3, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution(tt.args.list); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("solution() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
func Test_solutionB(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name:       "1",
			args:       args{[]int{1, 2, 3, 4, 5}},
			wantResult: []int{120, 60, 40, 30, 24},
		},
		{
			name:       "2",
			args:       args{[]int{3, 2, 1}},
			wantResult: []int{2, 3, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solutionB(tt.args.list); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("solutionB() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Benchmark_solution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	}
}
func Benchmark_solutionB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solutionB([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	}
}
