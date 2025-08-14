package p1382

import (
	"reflect"
	"testing"
)
import . "leetcode/stucture"

var null = NULL

func Test_helper(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{root: []int{1, null, 2, null, 3, null, 4, null, null}}, []int{2, 1, 3, null, null, null, 4}},
		{"Example 2", args{root: []int{2, 1, 3}}, []int{2, 1, 3}},
		{"Testcase 2", args{root: []int{14, 9, 16, 2, 13}}, []int{13, 2, 14, null, 9, null, 16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helper() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Day-Stout-Warren Algorithm

func Test_helperDSW(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{root: []int{1, null, 2, null, 3, null, 4, null, null}}, []int{3, 2, 4, 1}},
		{"Example 2", args{root: []int{2, 1, 3}}, []int{2, 1, 3}},
		{"Testcase 2", args{root: []int{14, 9, 16, 2, 13}}, []int{14, 9, 16, 2, 13}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helper() = %v, want %v", got, tt.want)
			}
		})
	}
}
