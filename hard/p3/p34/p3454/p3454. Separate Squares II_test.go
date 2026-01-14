package p3454

import "testing"

func Test_separateSquares(t *testing.T) {
	type args struct {
		squares [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Example 1", args{squares: [][]int{{0, 0, 1}, {2, 2, 1}}}, 1.00000},
		{"Example 2", args{squares: [][]int{{0, 0, 2}, {1, 1, 1}}}, 1.00000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := separateSquares(tt.args.squares); got != tt.want {
				t.Errorf("separateSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
