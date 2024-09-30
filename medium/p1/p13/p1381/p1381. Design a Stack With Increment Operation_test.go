package p1381

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		cmd   []string
		param [][]int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"Example 1", args{[]string{"CustomStack", "push", "push", "pop", "push", "push", "push", "increment", "increment", "pop", "pop", "pop", "pop"},
			[][]int{{3}, {1}, {2}, {}, {2}, {3}, {4}, {5, 100}, {2, 100}, {}, {}, {}, {}}}, []interface{}{nil, nil, nil, 2, nil, nil, nil, nil, nil, 103, 202, 201, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.cmd, tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
