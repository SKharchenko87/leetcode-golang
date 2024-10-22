package p2583

import (
	"leetcode/stucture"
	"testing"
)

const null = stucture.NULL

func Test_run(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{arr: []int{5, 8, 9, 2, 1, 3, 7, 4, 6}, k: 2}, 13},
		{"Example 2", args{arr: []int{1, 2, null, 3}, k: 1}, 3},
		{"TestCase 5", args{arr: []int{5, 8, 9, 2, 1, 3, 7}, k: 4}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
