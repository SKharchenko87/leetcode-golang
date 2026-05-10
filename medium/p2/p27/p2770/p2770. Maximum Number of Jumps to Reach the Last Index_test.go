package p2770

import "testing"

func Test_maximumJumps(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 3, 6, 4, 1, 2}, target: 2}, 3},
		{"Example 2", args{nums: []int{1, 3, 6, 4, 1, 2}, target: 3}, 5},
		{"Example 3", args{nums: []int{1, 3, 6, 4, 1, 2}, target: 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumJumps(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("maximumJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
