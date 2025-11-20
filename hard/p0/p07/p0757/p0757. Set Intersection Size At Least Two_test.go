package p0757

import "testing"

func Test_intersectionSizeTwo(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{intervals: [][]int{{1, 3}, {3, 7}, {8, 9}}}, 5},
		{"Example 2", args{intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}}, 3},
		{"Example 3", args{intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}}, 5},
		{"Testcase 49", args{intervals: [][]int{{6, 21}, {1, 15}, {15, 20}, {10, 21}, {0, 7}}}, 4},
		{"Testcase 101", args{intervals: [][]int{{1, 3}, {3, 7}, {5, 7}, {7, 8}}}, 5},
		{"Testcase 103", args{intervals: [][]int{{12, 19}, {18, 25}, {4, 6}, {19, 24}, {19, 22}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectionSizeTwo(tt.args.intervals); got != tt.want {
				t.Errorf("intersectionSizeTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
