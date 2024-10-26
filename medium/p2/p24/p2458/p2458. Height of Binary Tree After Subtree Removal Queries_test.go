package p2458

import (
	"leetcode/stucture"
	"reflect"
	"testing"
)

const null = stucture.NULL

func Test_run(t *testing.T) {
	type args struct {
		root    []int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{root: []int{1, 3, 4, 2, null, 6, 5, null, null, null, null, null, 7}, queries: []int{4}}, []int{2}},
		{"Example 2", args{root: []int{5, 8, 9, 2, 1, 3, 7, 4, 6}, queries: []int{3, 2, 4, 8}}, []int{3, 2, 3, 2}},
		{"TestCase 19", args{root: []int{1, null, 5, 3, null, 2, 4}, queries: []int{3, 5, 4, 2, 4}}, []int{1, 0, 3, 3, 3}},
		{"TestCase 29", args{root: []int{8, 6, 37, 3, 7, 33, 65, 1, 4, null, null, 29, 36, 51, 66, null, 2, null, 5, 26, 31, 35, null, 45, 58, null, null, null, null, null, null, 22, 28, 30, 32, 34, null, 44, 47, 55, 59, 21, 25, 27, null, null, null, null, null, null, null, 41, null, 46, 48, 54, 56, null, 62, 13, null, 24, null, null, null, 38, 42, null, null, null, 49, 53, null, null, 57, 60, 64, 11, 20, 23, null, null, 39, null, 43, null, 50, 52, null, null, null, null, 61, 63, null, 10, 12, 18, null, null, null, null, 40, null, null, null, null, null, null, null, null, null, null, 9, null, null, null, 16, 19, null, null, null, null, 15, 17, null, null, 14}, queries: []int{57, 7, 32, 55, 20, 25, 45, 34, 5, 19, 45, 30, 48, 1, 47, 32, 60, 31, 22, 15, 31}}, []int{12, 12, 12, 12, 10, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 9, 11, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.root, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
