package p3739

import "testing"

func Test_countMajoritySubarrays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{1, 2, 2, 3}, target: 2}, 5},
		{"Example 2", args{nums: []int{1, 1, 1, 1}, target: 1}, 10},
		{"Example 3", args{nums: []int{1, 2, 3}, target: 4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMajoritySubarrays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("countMajoritySubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
