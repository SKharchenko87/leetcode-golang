package p3025

import "testing"

func Test_numberOfPairs(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{points: [][]int{{1, 1}, {2, 2}, {3, 3}}}, 0},
		{"Example 2", args{points: [][]int{{6, 2}, {4, 4}, {2, 6}}}, 2},
		{"Example 3", args{points: [][]int{{3, 1}, {1, 3}, {1, 1}}}, 2},
		{"TestCase 272", args{points: [][]int{{1, 4}, {1, 5}, {0, 3}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPairs(tt.args.points); got != tt.want {
				t.Errorf("numberOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
