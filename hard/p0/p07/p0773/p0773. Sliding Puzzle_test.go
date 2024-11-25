package p0773

import "testing"

func Test_slidingPuzzle(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{board: [][]int{{1, 2, 3}, {4, 0, 5}}}, 1},
		{"Example 2", args{board: [][]int{{1, 2, 3}, {5, 4, 0}}}, -1},
		{"Example 3", args{board: [][]int{{4, 1, 2}, {5, 0, 3}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingPuzzle(tt.args.board); got != tt.want {
				t.Errorf("slidingPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_change(t *testing.T) {
	type args struct {
		n int
		i byte
		j byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 123456, i: 0, j: 5}, 623451},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.args.n, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("change() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getZeroIndex(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"Example 1", args{n: 123450}, 0},
		{"Example 1", args{n: 103452}, 4},
		{"Example 1", args{n: 13452}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getZeroIndex(tt.args.n); got != tt.want {
				t.Errorf("getZeroIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_x(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Example 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x()
		})
	}
}
