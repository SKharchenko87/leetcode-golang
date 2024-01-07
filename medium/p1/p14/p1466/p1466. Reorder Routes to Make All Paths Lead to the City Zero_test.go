package p1466

import "testing"

func Test_minReorder(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}}, 3},
		{"Case 2", args{5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}}, 2},
		{"Case 3", args{3, [][]int{{1, 0}, {2, 0}}}, 0},
		{"Case 4", args{5, [][]int{{101, 0}, {101, 102}, {103, 102}, {103, 104}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minReorder(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("minReorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
