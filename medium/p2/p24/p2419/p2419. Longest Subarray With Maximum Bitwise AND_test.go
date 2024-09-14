package p2419

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 3, 3, 2, 2}}, 2},
		{"Example 2", args{nums: []int{1, 2, 3, 4}}, 1},
		{"Example 3", args{nums: []int{1, 2, 3, 3, 2, 2, 3}}, 2},
		{"Example 4", args{nums: []int{1, 2, 3, 3, 2, 2, 3, 3, 3}}, 3},
		{"Example 5", args{nums: []int{96317, 96317, 96317, 96317, 96317, 96317, 96317, 96317, 96317, 279979}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
