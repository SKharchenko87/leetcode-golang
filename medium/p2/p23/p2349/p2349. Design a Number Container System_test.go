package p2349

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		commands []string
		params   [][]int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"Example 1", args{commands: []string{"NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"},
			params: [][]int{{}, {10}, {2, 10}, {1, 10}, {3, 10}, {5, 10}, {10}, {1, 20}, {10}}},
			[]interface{}{nil, -1, nil, nil, nil, nil, 1, nil, 2},
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
