package p1074

import "testing"

func Test_numSubmatrixSumTarget(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0}, 4},
		{"Case 2", args{[][]int{{1, -1}, {-1, 1}}, 0}, 5},
		{"Case 3", args{[][]int{{904}}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubmatrixSumTarget(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("numSubmatrixSumTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
