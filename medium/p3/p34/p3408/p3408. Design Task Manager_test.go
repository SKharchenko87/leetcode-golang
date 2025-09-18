package p3408

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		commands []string
		params   []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{"Example 1",
			args{
				[]string{"TaskManager", "add", "edit", "execTop", "rmv", "add", "execTop"},
				[]any{
					[][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}},
					[]int{4, 104, 5},
					[]int{102, 8},
					[]int{},
					[]int{101},
					[]int{5, 105, 15},
					[]int{},
				},
			},
			[]any{nil, nil, nil, 3, nil, nil, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.commands, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
