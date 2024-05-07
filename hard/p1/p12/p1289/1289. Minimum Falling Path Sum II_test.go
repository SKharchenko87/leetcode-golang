package p1289

import "testing"

func Test_minFallingPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, 13},
		{"Example 2", args{[][]int{{7}}}, 7},
		{"Example 3", args{[][]int{{-73, 61, 43, -48, -36}, {3, 30, 27, 57, 10}, {96, -76, 84, 59, -15}, {5, -49, 76, 31, -7}, {97, 91, 61, -46, 67}}}, -192},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
