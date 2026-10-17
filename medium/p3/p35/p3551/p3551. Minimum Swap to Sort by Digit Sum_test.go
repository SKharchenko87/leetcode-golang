package p3551

import "testing"

func Test_minSwap(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{37, 100}}, 1},
		{"Example 2", args{nums: []int{22, 14, 33, 7}}, 0},
		{"Example 3", args{nums: []int{18, 43, 34, 16}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwaps(tt.args.nums); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
