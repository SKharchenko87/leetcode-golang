package p3741

import "testing"

func Test_minimumDistance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 1, 1, 3}}, 6},
		{"Example 2", args{nums: []int{1, 1, 2, 3, 2, 1, 2}}, 8},
		{"Example 3", args{nums: []int{1}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDistance(tt.args.nums); got != tt.want {
				t.Errorf("minimumDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
