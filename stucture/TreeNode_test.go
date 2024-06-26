package stucture

import (
	"reflect"
	"testing"
)

func TestTreeNodeToSlice(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test 0", args{SliceToTreeNodeFullMinInt([]int{4, 1, 6, 0, 2, 5, 7, NULL, NULL, NULL, 3, NULL, NULL, NULL, 8, NULL, NULL, 9})}, []int{4, 1, 6, 0, 2, 5, 7, NULL, NULL, NULL, 3, NULL, NULL, NULL, 8, NULL, NULL, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreeNodeToSlice(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeNodeToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
