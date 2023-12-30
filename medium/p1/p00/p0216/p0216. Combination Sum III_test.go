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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
