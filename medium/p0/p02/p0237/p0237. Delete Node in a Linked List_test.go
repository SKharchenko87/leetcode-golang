package p0237

import (
	. "leetcode/stucture"
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		head *ListNode
		node int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example 1", args{SliceToListNode([]int{4, 5, 1, 9}), 5}, SliceToListNode([]int{4, 1, 9})},
		{"Example 2", args{SliceToListNode([]int{4, 5, 1, 9}), 1}, SliceToListNode([]int{4, 5, 9})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.head, tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
