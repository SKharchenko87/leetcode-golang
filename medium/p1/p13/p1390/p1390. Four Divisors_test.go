package p1390

import (
	"testing"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Run", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumFourDivisors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{21, 4, 7}}, 32},
		{"Example 2", args{nums: []int{21, 21}}, 64},
		{"Example 3", args{nums: []int{1, 2, 3, 4, 5}}, 0},
		{"TestCase 15", args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumFourDivisors(tt.args.nums); got != tt.want {
				t.Errorf("sumFourDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_debugMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			debugMap()
		})
	}
}
