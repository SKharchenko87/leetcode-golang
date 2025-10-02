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

func Test_maxPointsInsideSquare(t *testing.T) {
	type args struct {
		points [][]int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{points: [][]int{{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3}}, s: "abdca"}, 2},
		{"Example 2", args{points: [][]int{{1, 1}, {-2, -2}, {-2, 2}}, s: "abb"}, 1},
		{"Example 3", args{points: [][]int{{1, 1}, {-1, -1}, {2, -2}}, s: "ccd"}, 0},
		{"TestCase 436", args{points: [][]int{{-35, -3}, {17, 28}, {28, -28}, {25, -1}, {25, -16}, {1, -21}}, s: "ffcbea"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPointsInsideSquare(tt.args.points, tt.args.s); got != tt.want {
				t.Errorf("maxPointsInsideSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
