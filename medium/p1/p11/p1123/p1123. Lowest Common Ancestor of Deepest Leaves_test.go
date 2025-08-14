package p1123

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		root []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{"Example 1", args{root: []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}}, []any{2, 7, 4}},
		{"Example 2", args{root: []any{1}}, []any{1}},
		{"Example 3", args{root: []any{0, 1, 3, nil, 2}}, []any{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
