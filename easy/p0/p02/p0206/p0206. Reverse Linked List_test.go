package p0206

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		headArr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{1, 2, 3, 4, 5}}, []int{5, 4, 3, 2, 1}},
		{"Example 2", args{[]int{1, 2}}, []int{2, 1}},
		{"Example 3", args{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.headArr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
