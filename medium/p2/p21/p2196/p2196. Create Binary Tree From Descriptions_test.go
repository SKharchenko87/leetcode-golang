package p2196

import (
	"leetcode/stucture"
	"reflect"
	"testing"
)

const null = stucture.NULL

func Test_run(t *testing.T) {
	type args struct {
		descriptions [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}}, []int{50, 20, 80, 15, 17, 19}},
		{"Example 2", args{descriptions: [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}}}, []int{1, 2, null, null, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.descriptions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
