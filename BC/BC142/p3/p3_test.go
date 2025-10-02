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

func Test_maxScore(t *testing.T) {
	type args struct {
		n           int
		k           int
		stayScore   [][]int
		travelScore [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 2, k: 1, stayScore: [][]int{{2, 3}}, travelScore: [][]int{{0, 2}, {1, 0}}}, 3},
		{"Example 2", args{n: 3, k: 2, stayScore: [][]int{{3, 4, 2}, {2, 1, 2}}, travelScore: [][]int{{0, 2, 1}, {2, 0, 4}, {3, 2, 0}}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.n, tt.args.k, tt.args.stayScore, tt.args.travelScore); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
