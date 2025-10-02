package p2

import (
	"testing"
)

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestSubsequence(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{16, 6, 3}}, 3},
		{"Example 2", args{nums: []int{6, 5, 3, 4, 2, 1}}, 4},
		{"Example 3", args{nums: []int{10, 20, 10, 19, 10, 20}}, 5},
		{"Example 4", args{nums: []int{8, 5, 9, 3}}, 3},
		{"Example 4", args{nums: []int{4, 5, 2, 6, 9}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequence(tt.args.nums); got != tt.want {
				t.Errorf("longestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
