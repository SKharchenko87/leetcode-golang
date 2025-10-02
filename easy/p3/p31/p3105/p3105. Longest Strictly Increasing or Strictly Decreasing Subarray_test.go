package p3105

import "testing"

func Test_longestMonotonicSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 4, 3, 3, 2}}, 2},
		{"Example 2", args{nums: []int{3, 3, 3, 3}}, 1},
		{"Example 3", args{nums: []int{3, 2, 1}}, 3},
		{"TestCase 430", args{nums: []int{1, 1, 5}}, 2},
		{"TestCase 432", args{nums: []int{1, 10, 10}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMonotonicSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestMonotonicSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
