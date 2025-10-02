package p1261

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		commands []string
		params   [][]interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{

		{"Example 1", args{commands: []string{"FindElements", "find", "find"}, params: [][]interface{}{{-1, nil, -1}, {1}, {2}}}, []interface{}{nil, false, true}},
		{"Example 2", args{commands: []string{"FindElements", "find", "find", "find"}, params: [][]interface{}{{-1, -1, -1, -1, -1}, {1}, {3}, {5}}}, []interface{}{nil, true, true, false}},
		{"Example 3", args{commands: []string{"FindElements", "find", "find", "find", "find"}, params: [][]interface{}{{-1, nil, -1, -1, nil, -1}, {2}, {3}, {4}, {5}}}, []interface{}{nil, true, false, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.commands, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
