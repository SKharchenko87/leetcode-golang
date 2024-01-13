package p0450

import (
	. "leetcode/stucture"
	"reflect"
	"testing"
)

const null = NULL

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Case 1", args{SliceToTreeNodeFullMinInt([]int{5, 3, 6, 2, 4, null, 7}), 3}, SliceToTreeNodeFullMinInt([]int{5, 4, 6, 2, null, null, 7})},
		{"Case 2", args{SliceToTreeNodeFullMinInt([]int{5, 3, 6, 2, 4, null, 7}), 0}, SliceToTreeNodeFullMinInt([]int{5, 3, 6, 2, 4, null, 7})},
		{"Case 3", args{SliceToTreeNodeFullMinInt([]int{}), 0}, SliceToTreeNodeFullMinInt([]int{})},
		{"Case 3", args{SliceToTreeNodeFullMinInt([]int{0}), 0}, SliceToTreeNodeFullMinInt([]int{})},
		{"Case 13", args{SliceToTreeNodeFullMinInt([]int{2, 1}), 2}, SliceToTreeNodeFullMinInt([]int{1})},
		{"Case 7", args{SliceToTreeNodeFullMinInt([]int{5, 3, 6, 2, 4, null, 7}), 7}, SliceToTreeNodeFullMinInt([]int{5, 3, 6, 2, 4})},
		{"Case 7", args{SliceToTreeNodeFullMinInt([]int{50, 30, 70, null, 40, 60, 80}), 50}, SliceToTreeNodeFullMinInt([]int{5, 3, 6, 2, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
