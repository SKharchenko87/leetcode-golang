package p1352

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		commands []string
		param    []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"Example", args{
			[]string{"ProductOfNumbers", "add", "add", "add", "add", "add", "getProduct", "getProduct", "getProduct", "add", "getProduct"},
			[]interface{}{nil, 3, 0, 2, 5, 4, 2, 3, 4, 8, 2}},
			[]interface{}{nil, nil, nil, nil, nil, nil, 20, 40, 0, nil, 32},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.commands, tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
