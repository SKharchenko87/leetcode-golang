package p0938

import "testing"
import . "leetcode/stucture"

const null = NULL

func Test_rangeSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{SliceToTreeNode([]int{10, 5, 15, 3, 7, null, 18}, true, null), 7, 15}, 32},
		{"Case 2", args{SliceToTreeNode([]int{10, 5, 15, 3, 7, 13, 18, 1, null, 6}, true, null), 6, 10}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBST(tt.args.root, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
