package p0330

import "testing"

func Test_minPatches(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{1, 3}, 6}, 1},
		{"Example 2", args{[]int{1, 5, 10}, 20}, 2},
		{"Example 3", args{[]int{1, 2, 2}, 5}, 0},
		{"TestCase 79", args{[]int{2, 4, 14, 18, 20, 25, 25, 35, 73, 94}, 42}, 2},
		{"TestCase 82", args{[]int{1, 2, 31, 33}, 2147483647}, 28},
		{"TestCase 82", args{[]int{2, 9, 22, 28, 31, 38, 44, 44, 47, 52, 56, 61, 71, 77}, 42}, 3},
		{"TestCase 86", args{[]int{10, 30, 36, 42, 50, 76, 87, 88, 91, 92}, 96}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPatches(tt.args.nums, tt.args.n); got != tt.want {
				t.Errorf("minPatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
