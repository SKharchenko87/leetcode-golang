package p2780

import "testing"

func Test_minimumIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 2, 2}}, 2},
		{"Example 2", args{nums: []int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}}, 4},
		{"Example 3", args{nums: []int{3, 3, 3, 3, 7, 2, 2}}, -1},
		{"TestCase 839", args{nums: []int{1, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumIndex(tt.args.nums); got != tt.want {
				t.Errorf("minimumIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
