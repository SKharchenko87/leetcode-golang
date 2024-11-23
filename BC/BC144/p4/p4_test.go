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

func Test_maxCollectedFruits(t *testing.T) {
	type args struct {
		fruits [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{fruits: [][]int{{1, 2, 3, 4}, {5, 6, 8, 7}, {9, 10, 11, 12}, {13, 14, 15, 16}}}, 100},
		{"Example 2", args{fruits: [][]int{{1, 1}, {1, 1}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCollectedFruits(tt.args.fruits); got != tt.want {
				t.Errorf("maxCollectedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
