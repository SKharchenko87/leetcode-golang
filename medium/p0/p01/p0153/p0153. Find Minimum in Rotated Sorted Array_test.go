package p0153

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{3, 4, 5, 1, 2}}, 1},
		{"Example 2", args{nums: []int{4, 5, 6, 7, 0, 1, 2}}, 0},
		{"Example 3", args{nums: []int{11, 13, 15, 17}}, 11},
		{"TestCase 90", args{nums: []int{5, 1, 2, 3, 4}}, 1},
		{"TestCase 10", args{nums: []int{2, 3, 4, 5, 1}}, 1},
		{"TestCase 5", args{nums: []int{2, 1}}, 1},
		{"TestCase 7", args{nums: []int{3, 1, 2}}, 1},
		{"TestCase XX", args{nums: []int{7, 1, 2, 3, 4, 5, 6}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
