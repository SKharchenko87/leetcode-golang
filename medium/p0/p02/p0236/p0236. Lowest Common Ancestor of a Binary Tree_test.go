package main

import (
	. "leetcode/stucture"
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			// root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
			name: "anmes",
			args: args{
				SliceToTreeNodeFullMinInt([]int{3, 5, 1, 6, 2, NULL, 8, NULL, NULL, 7, 4}),
				SliceToTreeNodeFullMinInt([]int{5}),
				SliceToTreeNodeFullMinInt([]int{1})},
			want: SliceToTreeNodeFullMinInt([]int{3})},
		{
			// root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
			name: "anmes",
			args: args{
				SliceToTreeNodeFullMinInt([]int{3, 5, 1, 6, 2, NULL, 8, NULL, NULL, 7, 4}),
				SliceToTreeNodeFullMinInt([]int{5}),
				SliceToTreeNodeFullMinInt([]int{4})},
			want: SliceToTreeNodeFullMinInt([]int{5})},

		{
			//root = [1,2], p = 1, q = 2
			//output 1
			name: "anmes",
			args: args{
				SliceToTreeNodeFullMinInt([]int{1, 2}),
				SliceToTreeNodeFullMinInt([]int{1}),
				SliceToTreeNodeFullMinInt([]int{1})},
			want: SliceToTreeNodeFullMinInt([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
