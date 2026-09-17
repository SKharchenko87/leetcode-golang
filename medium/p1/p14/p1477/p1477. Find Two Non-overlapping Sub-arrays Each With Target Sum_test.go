package p1477

import "testing"

func Test_minSumOfLengths(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{arr: []int{3, 2, 2, 4, 3}, target: 3}, 2},
		{"Example 2", args{arr: []int{7, 3, 4, 7}, target: 7}, 2},
		{"Example 3", args{arr: []int{4, 3, 2, 6, 2, 3, 4}, target: 6}, -1},
		{"My 1", args{arr: []int{4, 3, 2, 4, 1, 1, 2, 3, 4}, target: 6}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSumOfLengths(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("minSumOfLengths() = %v, want %v", got, tt.want)
			}
		})
	}
}
