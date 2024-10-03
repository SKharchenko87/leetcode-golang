package p1331

import (
	"reflect"
	"testing"
)

func Test_arrayRankTransform(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{arr: []int{40, 10, 20, 30}}, []int{4, 1, 2, 3}},
		{"Example 2", args{arr: []int{100, 100, 100}}, []int{1, 1, 1}},
		{"Example 3", args{arr: []int{37, 12, 28, 9, 100, 56, 80, 5, 12}}, []int{5, 3, 4, 2, 8, 6, 7, 1, 3}},
		{"TestCase 28", args{arr: []int{-34, -47, 40, 3, -30, 29, -16, -45, 2, 25, 48, 23, 2, -43, 4, -25, 26, -19, 44, 3, -30, -43}}, []int{4, 1, 16, 10, 5, 15, 8, 2, 9, 13, 18, 12, 9, 3, 11, 6, 14, 7, 17, 10, 5, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayRankTransform(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayRankTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
