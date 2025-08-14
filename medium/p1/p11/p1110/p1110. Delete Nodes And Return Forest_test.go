package p1110

import (
	. "leetcode/stucture"
	"reflect"
	"testing"
)

const null = NULL

func Test_run(t *testing.T) {
	type args struct {
		slice     []int
		to_delete []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{[]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 5}}, [][]int{{1, 2, null, 4}, {6}, {7}}},
		{"Example 2", args{[]int{1, 2, 4, null, 3}, []int{3}}, [][]int{{1, 2, 4}}},
		{"Testcase 4", args{[]int{1, 2, 3, null, null, null, 4}, []int{2, 1}}, [][]int{{3, null, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.slice, tt.args.to_delete); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
