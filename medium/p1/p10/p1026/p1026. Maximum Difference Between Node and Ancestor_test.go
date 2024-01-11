package p1026

import "testing"
import . "leetcode/stucture"

const null = NULL

func Test_maxAncestorDiff(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{SliceToTreeNodeFullMinInt([]int{8, 3, 10, 1, 6, null, 14, null, null, 4, 7, 13})}, 7},
		{"Case 2", args{SliceToTreeNodeFullMinInt([]int{1, null, 2, null, 0, 3})}, 3},
		{"Case 17", args{SliceToTreeNodeFullMinInt([]int{2, 4, 3, 1, null, 0, 5, null, 6, null, null, null, 7})}, 5},
		{"Case 18", args{SliceToTreeNodeFullMinInt([]int{4, 7, 2, 3, 9, null, null, 12, null, 0, null, null, 6, null, null, null, 8, null, 11, null, 10, 1, null, 5})}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAncestorDiff(tt.args.root); got != tt.want {
				t.Errorf("maxAncestorDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
