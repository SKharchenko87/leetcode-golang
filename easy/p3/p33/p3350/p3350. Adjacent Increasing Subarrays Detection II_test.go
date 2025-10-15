package p3350

import "testing"

func Test_maxIncreasingSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}}, 3},
		{"Example 2", args{nums: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIncreasingSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("maxIncreasingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
