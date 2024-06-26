package p1642

import "testing"

func Test_furthestBuilding(t *testing.T) {
	type args struct {
		heights []int
		bricks  int
		ladders int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{4, 2, 7, 6, 9, 14, 12}, 5, 1}, 4},
		{"Example 2", args{[]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2}, 7},
		{"Example 3", args{[]int{14, 3, 19, 3}, 17, 0}, 3},
		{"Example 3", args{[]int{1, 13, 1, 1, 13, 5, 11, 11}, 10, 8}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := furthestBuilding(tt.args.heights, tt.args.bricks, tt.args.ladders); got != tt.want {
				t.Errorf("furthestBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}
