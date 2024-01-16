package p0380

import (
	"reflect"
	"testing"
)

const null = -1

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
		{"Case 1", args{[]string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"}, [][]int{{}, {1}, {2}, {2}, {}, {1}, {2}, {}}}, []interface{}{null, true, false, true, 2, true, false, 2}},
		{"Case 11", args{[]string{"RandomizedSet", "insert", "insert", "remove", "insert", "remove", "getRandom"}, [][]int{{}, {0}, {1}, {0}, {2}, {1}, {}}}, []interface{}{null, true, true, true, true, true, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.commands, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
