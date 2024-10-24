package p0951

import (
	"leetcode/stucture"
	"testing"
)

const null = stucture.NULL

func Test_run(t *testing.T) {
	type args struct {
		root1Slice []int
		root2Slice []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{[]int{1, 2, 3, 4, 5, 6, null, null, null, 7, 8}, []int{1, 3, 2, null, 6, 4, 5, null, null, null, null, 8, 7}}, true},
		{"Example 2", args{[]int{}, []int{}}, true},
		{"Example 3", args{[]int{}, []int{1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.root1Slice, tt.args.root2Slice); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
