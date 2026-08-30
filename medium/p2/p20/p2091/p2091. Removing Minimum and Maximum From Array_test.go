package p2091

import "testing"

func Test_minimumDeletions(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{2, 10, 7, 5, 4, 1, 8, 6}}, 5},
		{"Example 2", args{nums: []int{0, -4, 19, 1, 8, -2, -3, 5}}, 3},
		{"Example 3", args{nums: []int{101}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions(tt.args.nums); got != tt.want {
				t.Errorf("minimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
