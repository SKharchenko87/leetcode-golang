package p1

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

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{5, 2, 5, 4, 5}, 2}, 2},
		{"Example 2", args{[]int{2, 1, 2}, 2}, -1},
		{"Example 3", args{[]int{9, 7, 5, 3}, 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
