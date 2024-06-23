package p1438

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{8, 2, 4, 7}, limit: 4}, 2},
		{"Example 2", args{nums: []int{10, 1, 2, 4, 7, 2}, limit: 5}, 4},
		{"Example 3", args{nums: []int{4, 2, 2, 2, 4, 4, 2, 2}, limit: 0}, 3},
		{"Testcase 7", args{nums: []int{1, 2, 3, 4, 5, 2, 3, 4}, limit: 3}, 7},
		{"Testcase 24", args{nums: []int{1, 5, 6, 7, 8, 10, 6, 5, 6}, limit: 4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
