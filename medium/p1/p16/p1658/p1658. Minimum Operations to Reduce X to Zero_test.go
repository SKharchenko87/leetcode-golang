package p1658

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 1, 4, 2, 3}, x: 5}, 2},
		{"Example 2", args{nums: []int{5, 6, 7, 8, 9}, x: 4}, -1},
		{"Example 3", args{nums: []int{3, 2, 20, 1, 1, 3}, x: 10}, 5},
		{"TestCase 4", args{nums: []int{5, 2, 3, 1, 1}, x: 5}, 1},
		{"TestCase 5", args{nums: []int{1, 1}, x: 3}, -1},
		{"TestCase 95", args{nums: []int{1000, 1, 1, 2, 3}, x: 1004}, 3},
		{"TestCase 96", args{nums: []int{1, 10, 1}, x: 22}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
