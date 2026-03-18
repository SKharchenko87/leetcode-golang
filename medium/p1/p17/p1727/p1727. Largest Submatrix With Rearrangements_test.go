package p1727

import "testing"

func Test_largestSubmatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{matrix: [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}}, 4},
		{"Example 2", args{matrix: [][]int{{1, 0, 1, 0, 1}}}, 3},
		{"Example 3", args{matrix: [][]int{{1, 1, 0}, {1, 0, 1}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSubmatrix(tt.args.matrix); got != tt.want {
				t.Errorf("largestSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
