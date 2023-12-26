package p0700

import (
	. "leetcode/stucture"
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}

	l := &TreeNode{1, nil, nil}
	r := &TreeNode{3, nil, nil}
	l = &TreeNode{2, l, r}
	r = &TreeNode{7, nil, nil}
	root := &TreeNode{4, l, r}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"Case 1", args{root, 2}, l},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
