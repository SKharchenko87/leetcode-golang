package main

import (
	. "leetcode/stucture"
	"reflect"
	"testing"
)

func Test_deleteMiddle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example 1", args{head: ArrToList([]int{1, 3, 4, 7, 1, 2, 6})}, ArrToList([]int{1, 3, 4, 1, 2, 6})},
		{"Example 2", args{head: ArrToList([]int{1, 2, 3, 4})}, ArrToList([]int{1, 2, 4})},
		{"Example 3", args{head: ArrToList([]int{2, 1})}, ArrToList([]int{2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
