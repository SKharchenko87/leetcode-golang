package p0623

import (
	"reflect"
	"testing"
)
import . "leetcode/stucture"

const null = NULL

func Test_addOneRow(t *testing.T) {
	type args struct {
		root  *TreeNode
		val   int
		depth int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Example 1", args{SliceToTreeNodeFullMinInt([]int{4, 2, 6, 3, 1, 5}), 1, 2}, SliceToTreeNodeFullMinInt([]int{4, 1, 1, 2, null, null, 6, 3, 1, 5})},
		{"Example 2", args{SliceToTreeNodeFullMinInt([]int{4, 2, null, 3, 1}), 1, 3}, SliceToTreeNodeFullMinInt([]int{4, 2, null, 1, 1, 3, null, null, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOneRow(tt.args.root, tt.args.val, tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
