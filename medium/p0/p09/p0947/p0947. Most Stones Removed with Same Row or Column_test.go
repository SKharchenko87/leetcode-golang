package p0947

import "testing"

func Test_removeStones(t *testing.T) {
	type args struct {
		stones [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{stones: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}}, 5},
		{"Example 2", args{stones: [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}}, 3},
		{"Example 3", args{stones: [][]int{{0, 0}}}, 0},
		{"TestCase 14", args{stones: [][]int{{0, 1}, {1, 0}, {1, 1}}}, 2},
		{"TestCase 34", args{stones: [][]int{{0, 1}, {1, 2}, {0, 0}, {6, 7}, {7, 1}, {6, 6}, {2, 1}, {3, 7}, {0, 2}}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeStones(tt.args.stones); got != tt.want {
				t.Errorf("removeStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
