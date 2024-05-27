package p2597

import "testing"

func Test_beautifulSubsets(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{2, 4, 6}, 2}, 4},
		{"Example 2", args{[]int{1}, 1}, 1},
		{"Example 1*", args{[]int{2, 4, 6}, 1}, 7},
		{"Example 1*", args{[]int{2, 2, 6}, 1}, 7},
		{"Example 329", args{[]int{10, 4, 5, 7, 2, 1}, 3}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beautifulSubsets(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("beautifulSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
