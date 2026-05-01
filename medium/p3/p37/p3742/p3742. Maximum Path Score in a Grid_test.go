package p3742

import "testing"

func Test_maxPathScore(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathScore(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("maxPathScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
