package p2946

import "testing"

func Test_areSimilar(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, k: 4}, false},
		{"Example 2", args{mat: [][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}}, k: 2}, true},
		{"Example 3", args{mat: [][]int{{2, 2}, {2, 2}}, k: 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areSimilar(tt.args.mat, tt.args.k); got != tt.want {
				t.Errorf("areSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
