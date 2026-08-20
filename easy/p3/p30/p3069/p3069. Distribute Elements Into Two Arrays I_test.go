package p3069

import (
	"reflect"
	"testing"
)

func Test_resultArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{2, 1, 3}}, []int{2, 3, 1}},
		{"Example 2", args{nums: []int{5, 4, 3, 8}}, []int{5, 3, 4, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resultArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resultArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
