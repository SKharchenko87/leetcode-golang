package p1317

import (
	"reflect"
	"testing"
)

func Test_getNoZeroIntegers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{n: 2}, []int{1, 1}},
		{"Example 2", args{n: 11}, []int{2, 9}},
		{"TestCase 151", args{n: 10000}, []int{1, 9999}},
		{"TestCase 136", args{n: 160}, []int{1, 159}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNoZeroIntegers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNoZeroIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
