package p3

import "testing"

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

func Test_maxRemoval(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{2, 0, 2}, queries: [][]int{{0, 2}, {0, 2}, {1, 1}}}, 1},
		{"Example 2", args{nums: []int{1, 1, 1, 1}, queries: [][]int{{1, 3}, {0, 2}, {1, 3}, {1, 2}}}, 2},
		{"Example 3", args{nums: []int{1, 2, 3, 4}, queries: [][]int{{0, 3}}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRemoval(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("maxRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
