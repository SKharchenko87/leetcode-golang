package p3254

import (
	"reflect"
	"testing"
)

func Test_resultsArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{1, 2, 3, 4, 3, 2, 5}, k: 3}, []int{3, 4, -1, -1, -1}},
		{"Example 2", args{nums: []int{2, 2, 2, 2, 2}, k: 4}, []int{-1, -1}},
		{"Example 3", args{nums: []int{3, 2, 3, 2, 3, 2}, k: 2}, []int{-1, 3, -1, 3, -1}},
		{"Example 4", args{nums: []int{1, 4}, k: 1}, []int{1, 4}},
		{"Example 344", args{nums: []int{1, 3, 4}, k: 2}, []int{-1, 4}},
		{"Example 539", args{nums: []int{1, 30, 31, 32}, k: 3}, []int{-1, 32}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resultsArray(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resultsArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
