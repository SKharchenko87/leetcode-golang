package p2501

import "testing"

func Test_longestSquareStreak(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{4, 3, 6, 16, 8, 2}}, 3},
		{"Example 2", args{nums: []int{2, 3, 5, 6, 7}}, -1},
		{"TestCase 90", args{nums: []int{4, 16, 256, 65536}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSquareStreak(tt.args.nums); got != tt.want {
				t.Errorf("longestSquareStreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
