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
		// TODO: Add test cases.
		{
			// root = [5,3,6,2,4,null,7], key = 3
			name: "anmes",
			args: args{
				SliceToTreeNodeFullMinInt([]int{5, 3, 6, 2, 4, null, 7}),
				3},
			want: SliceToTreeNodeFullMinInt([]int{5, 4, 6, 2, null, null, 7})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
