package p1957

import "testing"

func Test_maxMatrixSum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{matrix: [][]int{{1, -1}, {-1, 1}}}, 4},
		{"Example 2", args{matrix: [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMatrixSum(tt.args.matrix); got != tt.want {
				t.Errorf("maxMatrixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
