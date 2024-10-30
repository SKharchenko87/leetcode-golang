package p1671

import "testing"

func Test_minimumMountainRemovals(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 3, 1}}, 0},
		{"Example 2", args{nums: []int{2, 1, 1, 5, 6, 2, 3, 1}}, 3},
		{"TestCase 15", args{nums: []int{23, 47, 63, 72, 81, 99, 88, 55, 21, 33, 32}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMountainRemovals(tt.args.nums); got != tt.want {
				t.Errorf("minimumMountainRemovals() = %v, want %v", got, tt.want)
			}
		})
	}
}
