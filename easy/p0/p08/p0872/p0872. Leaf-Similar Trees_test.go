package p0872

import "testing"
import . "leetcode/stucture"

const null = NULL

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{SliceToTreeNode([]int{3, 5, 1, 6, 2, 9, 8, null, null, 7, 4}, true, -1), SliceToTreeNode([]int{3, 5, 1, 6, 2, 9, 8, null, null, 7, 4}, true, -1)}, true},
		{"Case 2", args{SliceToTreeNode([]int{1, 2, 3}, true, -1), SliceToTreeNode([]int{1, 3, 2}, true, -1)}, false},
		{"Case 3", args{SliceToTreeNode([]int{3, 5, 1, 6, 2, 9, 8, null, null, 7, 4}, true, -1), SliceToTreeNode([]int{3, 5, 1, 6, 7, 4, 2, null, null, null, null, null, null, 9, 11, null, null, 8, 10}, true, -1)}, false},
		{"Case 4", args{SliceToTreeNode([]int{3, 5, 1, 6, 7, 4, 2, null, null, null, null, null, null, 9, 11, null, null, 8, 10}, true, -1), SliceToTreeNode([]int{3, 5, 1, 6, 2, 9, 8, null, null, 7, 4}, true, -1)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
