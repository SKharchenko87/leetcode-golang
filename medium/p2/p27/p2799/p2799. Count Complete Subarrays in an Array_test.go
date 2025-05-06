package p2799

import "testing"

func Test_countCompleteSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 3, 1, 2, 2}}, 4},
		{"Example 2", args{nums: []int{5, 5, 5, 5}}, 10},
		{"TestCase 232", args{nums: []int{1786, 1786, 1786, 114}}, 3},
		{"TestCase 412", args{nums: []int{1917, 1917, 608, 608, 1313, 751, 558, 1561, 608}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompleteSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("countCompleteSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
