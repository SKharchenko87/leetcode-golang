package p0543

import "testing"
import . "leetcode/stucture"

func Test_diameterOfBixnaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{SliceToTreeNodeFullMinInt([]int{1, 2, 3, 4, 5})}, 3},
		{"Example 2", args{SliceToTreeNodeFullMinInt([]int{1, 2})}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBixnaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
