package p3070

import "testing"

func Test_countSubmatrices(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{7, 6, 3}, {6, 6, 1}}, k: 18}, 4},
		{"Example 2", args{grid: [][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, k: 20}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubmatrices(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("countSubmatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
