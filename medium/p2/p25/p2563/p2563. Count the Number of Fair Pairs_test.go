package p2563

import "testing"

func Test_countFairPairs(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{0, 1, 7, 4, 4, 5}, lower: 3, upper: 6}, 6},
		{"Example 2", args{nums: []int{1, 7, 9, 2, 5}, lower: 11, upper: 11}, 1},
		{"TestCase 10", args{nums: []int{0, 0, 0, 0, 0, 0}, lower: 0, upper: 0}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countFairPairs(tt.args.nums, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("countFairPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
