package p2553

import (
	"reflect"
	"testing"
)

func Test_separateDigits(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{13, 25, 83, 77}}, []int{1, 3, 2, 5, 8, 3, 7, 7}},
		{"Example 2", args{nums: []int{7, 1, 3, 9}}, []int{7, 1, 3, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := separateDigits(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("separateDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
