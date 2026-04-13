package p1848

import "testing"

func Test_getMinDistance(t *testing.T) {
	type args struct {
		nums   []int
		target int
		start  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 3, 4, 5}, target: 5, start: 3}, 1},
		{"Example 2", args{nums: []int{1}, target: 1, start: 0}, 0},
		{"Example 3", args{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, target: 1, start: 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinDistance(tt.args.nums, tt.args.target, tt.args.start); got != tt.want {
				t.Errorf("getMinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
