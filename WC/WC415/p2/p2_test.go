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

func Test_maxScore(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{a: []int{3, 2, 5, 6}, b: []int{2, -6, 4, -5, -3, 2, -7}}, 26},
		{"Example 2", args{a: []int{-1, 4, 5, -2}, b: []int{-5, -1, -3, -2, -4}}, -1},
		{"Example 11", args{a: []int{9, 2, 1, 4}, b: []int{7, 2, 4, -6, 6, 9, 5, -1, -7, 9, 7, 8, 9, 5, 3, 7, 3, 9, 2, 0}}, 144},
		{"Example 437", args{a: []int{-2, 6, -1, 8}, b: []int{1, 6, -2, 0, 4, 6, -5, 4, -10, 4, 6, -10, 0, -9, -1, -9, 8, 9, 2}}, 138},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
