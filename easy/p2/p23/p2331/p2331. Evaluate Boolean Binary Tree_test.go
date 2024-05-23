package p2331

import (
	. "leetcode/stucture"
	"testing"
)

func Test_evaluateTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{SliceToTreeNodeFullMinInt([]int{2, 1, 3, NULL, NULL, 0, 1})}, true},
		{"Example 2", args{SliceToTreeNodeFullMinInt([]int{0})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateTree(tt.args.root); got != tt.want {
				t.Errorf("evaluateTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
