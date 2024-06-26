package p0216

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"Case 1", args{3, 7}, [][]int{{1, 2, 4}}},
		{"Case 1", args{3, 9}, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{"Case 1", args{4, 1}, [][]int{}},
		{"Case 1", args{9, 45}, [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}},
		{"Case 1", args{8, 36}, [][]int{{1, 2, 3, 4, 5, 6, 7, 8}}},
		{"Case 1", args{4, 24}, [][]int{{1, 6, 8, 9}, {2, 5, 8, 9}, {2, 6, 7, 9}, {3, 4, 8, 9}, {3, 5, 7, 9}, {3, 6, 7, 8}, {4, 5, 6, 9}, {4, 5, 7, 8}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
