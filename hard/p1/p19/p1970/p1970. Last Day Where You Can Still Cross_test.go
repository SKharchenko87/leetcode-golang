package p1970

import "testing"

func Test_latestDayToCross(t *testing.T) {
	type args struct {
		row   int
		col   int
		cells [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{row: 2, col: 2, cells: [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}}, 2},
		{"Example 2", args{row: 2, col: 2, cells: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}}, 1},
		{"Example 3", args{row: 3, col: 3, cells: [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := latestDayToCross(tt.args.row, tt.args.col, tt.args.cells); got != tt.want {
				t.Errorf("latestDayToCross() = %v, want %v", got, tt.want)
			}
		})
	}
}
