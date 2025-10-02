package p1937

import "testing"

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{points: [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}}, 9},
		{"Example 2", args{points: [][]int{{1, 5}, {2, 3}, {4, 2}}}, 11},
		{"TestCase 99", args{points: [][]int{{1, 5}, {3, 2}, {4, 2}}}, 11},
		{"TestCase 125", args{points: [][]int{{10}}}, 10},
		{"TestCase 135", args{points: [][]int{{1, 2, 3}}}, 3},
		{"Example 1", args{points: [][]int{{0, 3, 0, 4, 2}, {5, 4, 2, 4, 1}, {5, 0, 0, 5, 1}, {2, 0, 1, 0, 3}}}, 15},
		{"TestCase 139", args{points: [][]int{{5, 2, 1, 2}, {2, 1, 5, 2}, {5, 5, 5, 0}}}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
