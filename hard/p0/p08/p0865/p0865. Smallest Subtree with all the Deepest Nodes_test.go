package p0865

import (
	"reflect"
	"testing"
)

import . "leetcode/stucture"

func Test_subtreeWithAllDeepest(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Example 1", args{root: SliceToTreeNodeFullMinInt([]int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4})}, SliceToTreeNodeFullMinInt([]int{2, 7, 4})},
		{"Example 2", args{root: SliceToTreeNodeFullMinInt([]int{1})}, SliceToTreeNodeFullMinInt([]int{1})},
		{"Example 3", args{root: SliceToTreeNodeFullMinInt([]int{0, 1, 3, NULL, 2})}, SliceToTreeNodeFullMinInt([]int{2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtreeWithAllDeepest(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subtreeWithAllDeepest() = %v, want %v", got, tt.want)
			}
		})
	}
}
