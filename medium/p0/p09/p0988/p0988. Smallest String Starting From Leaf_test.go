package p0988

import "testing"
import . "leetcode/stucture"

const null = NULL

func Test_smallestFromLeaf(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{SliceToTreeNodeFullMinInt([]int{0, 1, 2, 3, 4, 3, 4})}, "dba"},
		{"Example 2", args{SliceToTreeNodeFullMinInt([]int{25, 1, 3, 1, 3, 0, 2})}, "adz"},
		{"Example 3", args{SliceToTreeNodeFullMinInt([]int{2, 2, 1, null, 1, 0, null, 0})}, "abc"},
		{"Example 45", args{SliceToTreeNodeFullMinInt([]int{0, null, 1})}, "ba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestFromLeaf(tt.args.root); got != tt.want {
				t.Errorf("smallestFromLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
