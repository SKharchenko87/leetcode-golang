package p1609

import "testing"
import . "leetcode/stucture"

func Test_isEvenOddTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{SliceToTreeNodeFullMinInt([]int{1, 10, 4, 3, NULL, 7, 9, 12, 8, 6, NULL, NULL, 2})}, true},
		{"Example 2", args{SliceToTreeNodeFullMinInt([]int{5, 4, 2, 3, 3, 7})}, false},
		{"Example 3", args{SliceToTreeNodeFullMinInt([]int{5, 9, 1, 3, 5, 7})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEvenOddTree(tt.args.root); got != tt.want {
				t.Errorf("isEvenOddTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
