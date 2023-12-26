package p0199

import (
	. "leetcode/stucture"
	"reflect"
	"testing"
)

const null = NULL

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "Case 1",
			args: args{SliceToTreeNodeFullMinInt([]int{1, 2, 3, null, 5, null, 4})},
			want: []int{1, 3, 4},
		},
		{
			name: "Case 2",
			args: args{SliceToTreeNodeFullMinInt([]int{1, null, 3})},
			want: []int{1, 3},
		},
		{
			name: "Case 3",
			args: args{SliceToTreeNodeFullMinInt([]int{})},
			want: []int{},
		},
		{
			name: "Case 4",
			args: args{SliceToTreeNodeFullMinInt([]int{1, 2, 3, null, 5})},
			want: []int{1, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
