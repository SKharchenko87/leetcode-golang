package p1886

import "testing"

func Test_findRotation(t *testing.T) {
	type args struct {
		mat    [][]int
		target [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{mat: [][]int{{0, 1}, {1, 0}}, target: [][]int{{1, 0}, {0, 1}}}, true},
		{"Example 2", args{mat: [][]int{{0, 1}, {1, 1}}, target: [][]int{{1, 0}, {0, 1}}}, false},
		{"Example 3", args{mat: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, target: [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}}, true},
		{"Example 110", args{mat: [][]int{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}}, target: [][]int{{1, 0, 0}, {0, 0, 1}, {0, 1, 0}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotation(tt.args.mat, tt.args.target); got != tt.want {
				t.Errorf("findRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
