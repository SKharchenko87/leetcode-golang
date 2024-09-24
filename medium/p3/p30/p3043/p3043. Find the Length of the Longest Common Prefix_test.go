package p3043

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{arr1: []int{1, 10, 100}, arr2: []int{1000}}, 3},
		{"Example 2", args{arr1: []int{1, 2, 3}, arr2: []int{4, 4, 4}}, 0},
		{"Example 2", args{arr1: []int{5435}, arr2: []int{543}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
