package p1304

import (
	"reflect"
	"testing"
)

func Test_sumZero(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{n: 5}, []int{-1, -2, 0, 2, 1}},
		{"Example 2", args{n: 3}, []int{-1, 0, 1}},
		{"Example 3", args{n: 1}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumZero(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
