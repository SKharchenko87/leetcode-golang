package p2616

import "testing"

func Test_minimizeMax(t *testing.T) {
	type args struct {
		nums []int
		p    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{10, 1, 2, 7, 1, 3}, p: 2}, 1},
		{"Example 2", args{nums: []int{4, 2, 1, 2}, p: 1}, 0},
		{"My 1", args{nums: []int{1, 2, 2, 3}, p: 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizeMax(tt.args.nums, tt.args.p); got != tt.want {
				t.Errorf("minimizeMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
