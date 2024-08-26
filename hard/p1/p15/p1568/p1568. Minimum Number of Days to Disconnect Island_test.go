package p1568

import "testing"

func Test_minDays(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{grid: [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}}, 2},
		{"Example 2", args{grid: [][]int{{1, 1}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDays(tt.args.grid); got != tt.want {
				t.Errorf("minDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
