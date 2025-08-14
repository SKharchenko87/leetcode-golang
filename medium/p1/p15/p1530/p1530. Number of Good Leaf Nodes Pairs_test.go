package p1530

import (
	"leetcode/stucture"
	"testing"
)

const null = stucture.NULL

func Test_run(t *testing.T) {
	type args struct {
		slice    []int
		distance int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{1, 2, 3, null, 4}, 3}, 1},
		{"Example 2", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, 2},
		{"Example 3", args{[]int{7, 1, 4, 6, null, 5, 3, null, null, null, null, null, 2}, 3}, 1},
		{"Testcase 46", args{[]int{15, 66, 55, 97, 60, 12, 56, null, 54, null, 49, null, 9, null, null, null, null, null, 90}, 5}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.slice, tt.args.distance); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
