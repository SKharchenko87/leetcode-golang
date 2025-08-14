package p2097

import (
	"reflect"
	"testing"
)

func Test_validArrangement(t *testing.T) {
	type args struct {
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{pairs: [][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}}, [][]int{{11, 9}, {9, 4}, {4, 5}, {5, 1}}},
		{"Example 2", args{pairs: [][]int{{1, 3}, {3, 2}, {2, 1}}}, [][]int{{1, 3}, {3, 2}, {2, 1}}},
		{"Example 3", args{pairs: [][]int{{1, 2}, {1, 3}, {2, 1}}}, [][]int{{1, 2}, {2, 1}, {1, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validArrangement(tt.args.pairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validArrangement() = %v, want %v", got, tt.want)
			}
		})
	}
}
