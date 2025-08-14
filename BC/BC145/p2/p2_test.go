package p2

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

func Test_findMinimumTime(t *testing.T) {
	type args struct {
		strength []int
		K        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{strength: []int{3, 4, 1}, K: 1}, 4},
		{"Example 2", args{strength: []int{2, 5, 4}, K: 2}, 5},
		{"Example 2", args{strength: []int{3}, K: 8}, 3},
		{"Example 2", args{strength: []int{7, 3, 6, 18, 22, 50}, K: 4}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinimumTime(tt.args.strength, tt.args.K); got != tt.want {
				t.Errorf("findMinimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
