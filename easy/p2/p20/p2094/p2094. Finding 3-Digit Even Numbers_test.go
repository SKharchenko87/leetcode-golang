package p2094

import (
	"reflect"
	"testing"
)

func Test_findEvenNumbers(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{digits: []int{2, 1, 3, 0}}, []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320}},
		{"Example 2", args{digits: []int{2, 2, 8, 8, 2}}, []int{222, 228, 282, 288, 822, 828, 882}},
		{"Example 3", args{digits: []int{3, 7, 5}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEvenNumbers(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findEvenNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
