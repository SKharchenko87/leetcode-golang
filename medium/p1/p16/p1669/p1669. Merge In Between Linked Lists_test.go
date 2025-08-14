package p1669

import (
	. "leetcode/stucture"
	"reflect"
	"testing"
)

func Test_mergeInBetween(t *testing.T) {
	type args struct {
		list1 *ListNode
		a     int
		b     int
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example 1", args{
			SliceToListNode([]int{10, 1, 13, 6, 9, 5}),
			3, 4,
			SliceToListNode([]int{1000000, 1000001, 1000002})},
			SliceToListNode([]int{10, 1, 13, 1000000, 1000001, 1000002, 5})},
		{"Example 2", args{
			SliceToListNode([]int{0, 1, 2, 3, 4, 5, 6}),
			2,
			5,
			SliceToListNode([]int{1000000, 1000001, 1000002, 1000003, 1000004})},
			SliceToListNode([]int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeInBetween(tt.args.list1, tt.args.a, tt.args.b, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeInBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
