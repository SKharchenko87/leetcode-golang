package p2054

import "testing"

func Test_maxTwoEvents(t *testing.T) {
	type args struct {
		events [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{events: [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}}, 4},
		{"Example 2", args{events: [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}}, 5},
		{"Example 3", args{events: [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}}, 8},
		{"TestCase 14", args{events: [][]int{{10, 83, 53}, {63, 87, 45}, {97, 100, 32}, {51, 61, 16}}}, 85},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTwoEvents(tt.args.events); got != tt.want {
				t.Errorf("maxTwoEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
