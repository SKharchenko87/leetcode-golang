package p0450

import (
	"fmt"
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
	tn10 := &TreeNode{10, nil, nil}
	tn9 := &TreeNode{9, nil, tn10}
	tn5 := &TreeNode{5, nil, nil}
	tn3 := &TreeNode{3, nil, nil}
	tn4 := &TreeNode{4, tn3, tn5}
	tn7 := &TreeNode{7, tn4, nil}
	tn1 := &TreeNode{1, nil, nil}
	tn2 := &TreeNode{2, tn1, tn7}
	tn8 := &TreeNode{8, tn2, tn9}

	wtn10 := &TreeNode{10, nil, nil}
	wtn9 := &TreeNode{9, nil, wtn10}
	wtn5 := &TreeNode{5, nil, nil}
	wtn4 := &TreeNode{4, nil, wtn5}
	wtn7 := &TreeNode{7, wtn4, nil}
	wtn1 := &TreeNode{1, nil, nil}
	wtn3 := &TreeNode{3, wtn1, wtn7}
	//wtn2 := &TreeNode{2, wtn1, wtn7}
	wtn8 := &TreeNode{8, wtn3, wtn9}
	//tn6 := &TreeNode{6, nil, nil}

	c4tn5 := &TreeNode{5, nil, nil}
	c4tn4 := &TreeNode{4, nil, c4tn5}
	c4tn1 := &TreeNode{1, nil, nil}
	c4tn7 := &TreeNode{7, c4tn4, nil}
	c4tn10 := &TreeNode{10, nil, nil}
	c4tn2 := &TreeNode{2, c4tn1, c4tn7}
	c4tn9 := &TreeNode{9, nil, c4tn10}
	c4tn8 := &TreeNode{8, c4tn2, c4tn9}

	c4wtn5 := &TreeNode{5, nil, nil}
	c4wtn1 := &TreeNode{1, nil, nil}
	c4wtn7 := &TreeNode{7, c4wtn5, nil}
	c4wtn10 := &TreeNode{10, nil, nil}
	c4wtn4 := &TreeNode{4, c4wtn1, c4wtn7}
	c4wtn9 := &TreeNode{9, nil, c4wtn10}
	//c4wtn3 := &TreeNode{3, c4wtn1, c4wtn7}
	//c4wtn2 := &TreeNode{2, c4wtn1, c4wtn7}
	c4wtn8 := &TreeNode{8, c4wtn4, c4wtn9}

	c0tn0 := &TreeNode{0, nil, nil}
	//c0wtn0 := nil

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			// root = [5,3,6,2,4,null,7], key = 3
			name: "CAse ",
			args: args{
				tn8,
				2},
			want: wtn8},
		{
			// root = [5,3,6,2,4,null,7], key = 3
			name: "Case 4",
			args: args{
				c4tn8,
				2},
			want: c4wtn8},
		{
			// root = [5,3,6,2,4,null,7], key = 3
			name: "Case 0",
			args: args{
				c0tn0,
				2},
			want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteNode(tt.args.root, tt.args.key)
			fmt.Println(ToMatrix(got))
			fmt.Println(ToMatrix(tt.want))
			fmt.Println(got)
			if !reflect.DeepEqual(ToMatrix(got), ToMatrix(tt.want)) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
