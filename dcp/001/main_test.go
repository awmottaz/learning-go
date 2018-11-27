// Daily Coding Problem 001
//
// Given a list of numbers and a number k, return whether any two numbers
// from the list add up to k.
//
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
//
// Bonus: Can you do this in one pass?

package main

import "testing"

func Test_finalSolution(t *testing.T) {
	type args struct {
		list   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{[]int{10, 15, 3, 7}, 17},
			want: true,
		},
		{
			name: "2",
			args: args{[]int{10, 15, 3, 7}, 13},
			want: true,
		},
		{
			name: "3",
			args: args{[]int{-5, 8, 12, 5, -11}, 8},
			want: false,
		},
		{
			name: "4",
			args: args{[]int{-5, 8, 12, 5, -11}, 3},
			want: true,
		},
		{
			name: "5",
			args: args{[]int{-5, 8, 12, 5, -11}, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalSolution(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("solutionOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_solutionOne(t *testing.T) {
	type args struct {
		list   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{[]int{10, 15, 3, 7}, 17},
			want: true,
		},
		{
			name: "2",
			args: args{[]int{10, 15, 3, 7}, 13},
			want: true,
		},
		{
			name: "3",
			args: args{[]int{-5, 8, 12, 5, -11}, 8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solutionOne(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("solutionOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_solutionTwo(t *testing.T) {
	type args struct {
		list   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{[]int{3, 7, 10, 15}, 17},
			want: true,
		},
		{
			name: "2",
			args: args{[]int{3, 7, 10, 15}, 13},
			want: true,
		},
		{
			name: "3",
			args: args{[]int{-11, -5, 5, 8, 12}, 8},
			want: false,
		},
		{
			name: "3",
			args: args{[]int{-11, -5, 5, 8, 12}, 3},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solutionOne(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("solutionOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
