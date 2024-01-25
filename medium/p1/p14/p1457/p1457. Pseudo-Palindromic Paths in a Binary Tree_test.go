package p1457

import "testing"
import . "leetcode/stucture"

const null = NULL

func Test_pseudoPalindromicPaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{SliceToTreeNodeFullMinInt([]int{2, 3, 1, 3, 1, null, 1})}, 2},
		{"Case 2", args{SliceToTreeNodeFullMinInt([]int{2, 1, 1, 1, 3, null, null, null, null, null, 1})}, 1},
		{"Case 3", args{SliceToTreeNodeFullMinInt([]int{9})}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoPalindromicPaths(tt.args.root); got != tt.want {
				t.Errorf("pseudoPalindromicPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
