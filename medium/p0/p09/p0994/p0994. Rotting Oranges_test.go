package p0994

import "testing"

func Test_orangesRotting(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}}, 4},
		{"Case 2", args{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}}, -1},
		{"Case 3", args{[][]int{{0, 2}}}, 0},
		{"Case 4", args{[][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}}}, 2},
		{"Case 200", args{[][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orangesRotting(tt.args.grid); got != tt.want {
				t.Errorf("orangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}
