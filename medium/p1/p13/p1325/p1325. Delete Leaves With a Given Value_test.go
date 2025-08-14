package p1325

import (
	. "leetcode/stucture"
	"reflect"
	"testing"
)

const null = NULL

func Test_removeLeafNodes(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Example 1", args{SliceToTreeNodeFullMinInt([]int{1, 2, 3, 2, null, 2, 4}), 2}, SliceToTreeNodeFullMinInt([]int{1, null, 3, null, 4})},
		{"Example 2", args{SliceToTreeNodeFullMinInt([]int{1, 3, 3, 3, 2}), 3}, SliceToTreeNodeFullMinInt([]int{1, 3, null, null, 2})},
		{"Example 3", args{SliceToTreeNodeFullMinInt([]int{1, 2, null, 2, null, 2}), 2}, SliceToTreeNodeFullMinInt([]int{1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeLeafNodes(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeLeafNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
