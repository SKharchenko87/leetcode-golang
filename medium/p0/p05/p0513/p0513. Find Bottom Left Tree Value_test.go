package p0513

import "testing"
import . "leetcode/stucture"

func Test_findBottomLeftValue(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{SliceToTreeNodeFullMinInt([]int{2, 1, 3})}, 1},
		{"Example 2", args{SliceToTreeNodeFullMinInt([]int{1, 2, 3, 4, NULL, 5, 6, NULL, NULL, 7})}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
