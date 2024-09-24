package p0241

import (
	"reflect"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{expression: "2-1-1"}, []int{0, 2}},
		{"Example 2", args{expression: "2*3-4*5"}, []int{-34, -14, -10, -10, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diffWaysToCompute(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
