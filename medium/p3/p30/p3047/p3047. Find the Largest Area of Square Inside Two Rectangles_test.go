package p3047

import "testing"

func Test_largestSquareArea(t *testing.T) {
	type args struct {
		bottomLeft [][]int
		topRight   [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{bottomLeft: [][]int{{1, 1}, {2, 2}, {3, 1}}, topRight: [][]int{{3, 3}, {4, 4}, {6, 6}}}, 1},
		{"Example 2", args{bottomLeft: [][]int{{1, 1}, {1, 3}, {1, 5}}, topRight: [][]int{{5, 5}, {5, 7}, {5, 9}}}, 4},
		{"Example 3", args{bottomLeft: [][]int{{1, 1}, {2, 2}, {1, 2}}, topRight: [][]int{{3, 3}, {4, 4}, {3, 4}}}, 1},
		{"Example 4", args{bottomLeft: [][]int{{1, 1}, {3, 3}, {3, 1}}, topRight: [][]int{{2, 2}, {4, 4}, {4, 2}}}, 0},
		{"TestCase 514", args{bottomLeft: [][]int{{3, 2}, {3, 1}, {6, 1}}, topRight: [][]int{{8, 3}, {4, 8}, {9, 10}}}, 1},
		{"TestCase 520", args{bottomLeft: [][]int{{2, 1}, {4, 15}, {17, 6}, {19, 5}}, topRight: [][]int{{3, 17}, {10, 23}, {23, 11}, {22, 13}}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSquareArea(tt.args.bottomLeft, tt.args.topRight); got != tt.want {
				t.Errorf("largestSquareArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
