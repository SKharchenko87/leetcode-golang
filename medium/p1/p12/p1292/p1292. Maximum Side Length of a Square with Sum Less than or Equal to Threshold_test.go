package p1292

import "testing"

func Test_maxSideLength(t *testing.T) {
	type args struct {
		mat       [][]int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{mat: [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}, threshold: 4}, 2},
		{"Example 2", args{mat: [][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}}, threshold: 1}, 0},
		{"TestCase 6", args{mat: [][]int{{1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}}, threshold: 6}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSideLength(tt.args.mat, tt.args.threshold); got != tt.want {
				t.Errorf("maxSideLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
