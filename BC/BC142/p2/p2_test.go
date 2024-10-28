package p2

import (
	"reflect"
	"testing"
)

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSubtreeSizes(t *testing.T) {
	type args struct {
		parent []int
		s      string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{parent: []int{-1, 0, 0, 1, 1, 1}, s: "abaabc"}, []int{6, 3, 1, 1, 1, 1}},
		{"Example 2", args{parent: []int{-1, 0, 4, 0, 1}, s: "abbba"}, []int{5, 2, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubtreeSizes(tt.args.parent, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubtreeSizes() = %v, want %v", got, tt.want)
			}
		})
	}
}
