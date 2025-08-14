package p4

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

func Test_subsequencesWithMiddleMode(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 1, 1, 1, 1, 1}}, 6},
		{"Example 2", args{nums: []int{1, 2, 2, 3, 3, 4}}, 4},
		{"Example 3", args{nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsequencesWithMiddleMode(tt.args.nums); got != tt.want {
				t.Errorf("subsequencesWithMiddleMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
