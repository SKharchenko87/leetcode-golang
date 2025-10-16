package p2598

import "testing"

func Test_findSmallestInteger(t *testing.T) {
	type args struct {
		nums  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, -10, 7, 13, 6, 8}, value: 5}, 4},
		{"Example 2", args{nums: []int{1, -10, 7, 13, 6, 8}, value: 7}, 2},
		{"TestCase 807", args{nums: []int{3, 0, 3, 2, 4, 2, 1, 1, 0, 4}, value: 5}, 10},
		{"TestCase 1032", args{nums: []int{0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1}, value: 2}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSmallestInteger(tt.args.nums, tt.args.value); got != tt.want {
				t.Errorf("findSmallestInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
