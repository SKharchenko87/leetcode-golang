package p2579

import "testing"

func Test_coloredCells(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{n: 1}, 1},
		{"Example 2", args{n: 2}, 5},
		{"Example 3", args{n: 3}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coloredCells(tt.args.n); got != tt.want {
				t.Errorf("coloredCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
