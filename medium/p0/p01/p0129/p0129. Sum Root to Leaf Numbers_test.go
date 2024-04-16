package p0129

import "testing"
import . "leetcode/stucture"

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{SliceToTreeNodeFullMinInt([]int{1, 2, 3})}, 25},
		{"Example 2", args{SliceToTreeNodeFullMinInt([]int{4, 9, 0, 5, 1})}, 1026},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
