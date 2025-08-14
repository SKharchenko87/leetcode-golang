package p2444

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		minK int
		maxK int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{1, 3, 5, 2, 7, 5}, minK: 1, maxK: 5}, 2},
		{"Example 2", args{nums: []int{1, 1, 1, 1}, minK: 1, maxK: 1}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums, tt.args.minK, tt.args.maxK); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
