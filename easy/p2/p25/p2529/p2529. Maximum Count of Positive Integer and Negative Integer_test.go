package p2529

import "testing"

func Test_maximumCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{-2, -1, -1, 1, 2, 3}}, 3},
		{"Example 2", args{nums: []int{-3, -2, -1, 0, 0, 1, 2}}, 3},
		{"Example 3", args{nums: []int{5, 20, 66, 1314}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCount(tt.args.nums); got != tt.want {
				t.Errorf("maximumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
