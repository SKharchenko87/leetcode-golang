package p2130

import "testing"
import . "leetcode/stucture"

func Test_pairSum(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{head: ArrToList([]int{5, 4, 2, 1})}, 6},
		{"Example 2", args{head: ArrToList([]int{4, 2, 2, 3})}, 7},
		{"Example 3", args{head: ArrToList([]int{1, 100000})}, 100001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairSum(tt.args.head); got != tt.want {
				t.Errorf("pairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
