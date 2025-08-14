package p2096

import (
	"leetcode/stucture"
	"testing"
)

const null = stucture.NULL

func Test_run(t *testing.T) {
	type args struct {
		slice      []int
		startValue int
		destValue  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{[]int{5, 1, 2, 3, null, 6, 4}, 3, 6}, "UURL"},
		{"Example 1", args{[]int{5, 1, 2, 3, 8, 6, 4}, 3, 6}, "UURL"},
		{"Example 1", args{[]int{5, 1, 2, 3, null, 6, 4}, 6, 3}, "UULL"},
		{"Example 1", args{[]int{5, 1, 2, 3, null, 6, 4}, 3, 4}, "UURR"},
		{"Example 2", args{[]int{2, 1}, 2, 1}, "L"},
		{"Example 2", args{[]int{1, 2}, 2, 1}, "U"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.slice, tt.args.startValue, tt.args.destValue); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
