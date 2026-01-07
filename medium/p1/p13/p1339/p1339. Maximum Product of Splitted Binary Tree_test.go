package p1339

import "testing"
import . "leetcode/stucture"

func Test_maxProduct(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{root: SliceToTreeNodeFullMinInt([]int{1, 2, 3, 4, 5, 6})}, 110},
		{"Example 2", args{root: SliceToTreeNodeFullMinInt([]int{1, NULL, 2, 3, 4, NULL, NULL, 5, 6})}, 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.root); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
