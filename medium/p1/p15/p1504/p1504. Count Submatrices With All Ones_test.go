package p1504

import "testing"

func Test_numSubmat(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{mat: [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}}, 13},
		{"Example 2", args{mat: [][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}}}, 24},
		{"TestCase 13", args{mat: [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}}, 5},
		{"TestCase 16", args{mat: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 1, 1}, {1, 1, 0}, {0, 1, 1}}}, 12},
		{"My 1", args{mat: [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 1}}}, 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubmat(tt.args.mat); got != tt.want {
				t.Errorf("numSubmat() = %v, want %v", got, tt.want)
			}
		})
	}
}
