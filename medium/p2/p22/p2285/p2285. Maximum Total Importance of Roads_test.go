package p2285

import "testing"

func Test_maximumImportance(t *testing.T) {
	type args struct {
		n     int
		roads [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{n: 5, roads: [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}}, 43},
		{"Example 2", args{n: 5, roads: [][]int{{0, 3}, {2, 4}, {1, 3}}}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumImportance(tt.args.n, tt.args.roads); got != tt.want {
				t.Errorf("maximumImportance() = %v, want %v", got, tt.want)
			}
		})
	}
}
