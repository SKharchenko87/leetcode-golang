package p1865

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
		{"Example 1", args{[]string{"FindSumPairs", "count", "add", "count", "count", "add", "add", "count"}, []any{[][]int{{1, 1, 2, 2, 2, 3}, {1, 4, 5, 2, 5, 4}}, []int{7}, []int{3, 2}, []int{8}, []int{4}, []int{0, 1}, []int{1, 1}, []int{7}}}, []any{nil, 8, nil, 2, 1, nil, nil, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.commands, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
