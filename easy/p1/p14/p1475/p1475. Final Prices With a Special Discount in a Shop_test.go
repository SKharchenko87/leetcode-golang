package p1475

import (
	"reflect"
	"testing"
)

func Test_finalPrices(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{prices: []int{8, 4, 6, 2, 3}}, []int{4, 2, 4, 2, 3}},
		{"Example 2", args{prices: []int{1, 2, 3, 4, 5}}, []int{1, 2, 3, 4, 5}},
		{"Example 3", args{prices: []int{10, 1, 1, 6}}, []int{9, 0, 1, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalPrices(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finalPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
