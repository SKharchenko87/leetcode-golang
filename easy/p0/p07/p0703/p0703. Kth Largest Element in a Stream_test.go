package p0703

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		commands   []string
		k          int
		initParams []int
		params     []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{commands: []string{"add", "add", "add", "add", "add"}, k: 3, initParams: []int{4, 5, 8, 2}, params: []int{3, 5, 10, 9, 4}}, []int{4, 5, 5, 8, 8}},
		{"TestCase 6", args{commands: []string{"add", "add", "add", "add", "add"}, k: 2, initParams: []int{0}, params: []int{-1, 1, -2, -4, 3}}, []int{-1, 0, 0, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.commands, tt.args.k, tt.args.initParams, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
