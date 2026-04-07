package p2069

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		commands []string
		values   [][]int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"Example 1", args{[]string{"Robot", "step", "step", "getPos", "getDir", "step", "step", "step", "getPos", "getDir"}, [][]int{{6, 3}, {2}, {2}, {}, {}, {2}, {1}, {4}, {}, {}}}, []interface{}{nil, nil, nil, []int{4, 0}, "East", nil, nil, nil, []int{1, 2}, "West"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.commands, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
