package p3197

import "testing"

func Test_minimumSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{1, 0, 1}, {1, 1, 1}}}, 5},
		{"Example 2", args{grid: [][]int{{1, 0, 1, 0}, {0, 1, 0, 1}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSum(tt.args.grid); got != tt.want {
				t.Errorf("minimumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
