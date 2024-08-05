package p2134

import "testing"

func Test_minSwaps(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{0, 1, 0, 1, 1, 0, 0}}, 1},
		{"Example 2", args{nums: []int{0, 1, 1, 1, 0, 0, 1, 1, 0}}, 2},
		{"Example 3", args{nums: []int{1, 1, 0, 0, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwaps(tt.args.nums); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
