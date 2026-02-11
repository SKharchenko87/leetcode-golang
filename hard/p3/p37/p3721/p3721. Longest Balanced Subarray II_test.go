package p3721

import "testing"

func Test_longestBalanced(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{2, 5, 4, 3}}, 4},
		{"Example 2", args{nums: []int{3, 2, 2, 5, 4}}, 5},
		{"Example 3", args{nums: []int{1, 2, 3, 2}}, 3},
		{"TestCase 514", args{nums: []int{30, 27, 25}}, 2},
		{"TestCase 582", args{nums: []int{11, 34, 44, 22}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestBalanced(tt.args.nums); got != tt.want {
				t.Errorf("longestBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
