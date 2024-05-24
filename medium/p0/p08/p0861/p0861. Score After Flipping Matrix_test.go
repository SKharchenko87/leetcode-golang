package p0861

import "testing"

func Test_matrixScore(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}}, 39},
		{"Example 2", args{[][]int{{0}}}, 1},
		{"Example 2", args{[][]int{{1, 1, 1, 1, 0}, {0, 1, 0, 1, 1}, {1, 0, 0, 1, 1}}}, 78},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixScore(tt.args.grid); got != tt.want {
				t.Errorf("matrixScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
