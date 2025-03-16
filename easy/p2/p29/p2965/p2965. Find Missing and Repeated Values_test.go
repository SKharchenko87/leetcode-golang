package p2965

import (
	"reflect"
	"testing"
)

func Test_findMissingAndRepeatedValues(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{grid: [][]int{{1, 3}, {2, 2}}}, []int{2, 4}},
		{"Example 2", args{grid: [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}}, []int{9, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingAndRepeatedValues(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingAndRepeatedValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
