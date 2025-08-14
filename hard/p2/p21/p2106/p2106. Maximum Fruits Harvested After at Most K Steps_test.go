package p2106

import "testing"

func Test_maxTotalFruits(t *testing.T) {
	type args struct {
		fruits   [][]int
		startPos int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{fruits: [][]int{{2, 8}, {6, 3}, {8, 6}}, startPos: 5, k: 4}, 9},
		{"Example 2", args{fruits: [][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}, startPos: 5, k: 4}, 14},
		{"Example 3", args{fruits: [][]int{{0, 3}, {6, 4}, {8, 5}}, startPos: 3, k: 2}, 0},
		{"TestCase 81", args{fruits: [][]int{{200000, 10000}}, startPos: 0, k: 200000}, 10000},
		{"TestCase 82", args{fruits: [][]int{{200000, 10000}}, startPos: 200000, k: 200000}, 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTotalFruits(tt.args.fruits, tt.args.startPos, tt.args.k); got != tt.want {
				t.Errorf("maxTotalFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
