package p2099

import (
	"reflect"
	"testing"
)

func Test_maxSubsequence(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{2, 1, 3, 3}, k: 2}, []int{3, 3}},
		{"Example 2", args{nums: []int{-1, -2, 3, 4}, k: 3}, []int{-1, 3, 4}},
		{"Example 3", args{nums: []int{3, 4, 3, 3}, k: 2}, []int{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubsequence(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
