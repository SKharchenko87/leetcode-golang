package p0931

import "testing"

func Test_minFallingPathSum(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}}, 13},
		{"Case 2", args{[][]int{{-19, 57}, {-40, -5}}}, -59},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSum(tt.args.matrix); got != tt.want {
				t.Errorf("minFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
