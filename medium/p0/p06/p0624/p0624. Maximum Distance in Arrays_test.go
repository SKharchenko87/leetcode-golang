package p0624

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{arrays: [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}}, 4},
		{"Example 2", args{arrays: [][]int{{1}, {1}}}, 0},
		{"TestCase 102", args{arrays: [][]int{{-6, -3, -1, 1, 2, 2, 2}, {-10, -8, -6, -2, 4}, {-2}, {-8, -4, -3, -3, -2, -1, 1, 2, 3}, {-8, -6, -5, -4, -2, -2, 2, 4}}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.arrays); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
