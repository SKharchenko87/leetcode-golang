package p3550

import "testing"

func Test_smallestIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 3, 2}}, 2},
		{"Example 2", args{nums: []int{1, 10, 11}}, 1},
		{"Example 3", args{nums: []int{1, 2, 3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestIndex(tt.args.nums); got != tt.want {
				t.Errorf("smallestIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
