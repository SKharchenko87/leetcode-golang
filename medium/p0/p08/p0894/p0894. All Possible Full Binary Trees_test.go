package p0894

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]any
	}{
		{"Example 1", args{n: 7}, [][]any{{0, 0, 0, nil, nil, 0, 0, nil, nil, 0, 0}, {0, 0, 0, nil, nil, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, nil, nil, nil, nil, 0, 0}, {0, 0, 0, 0, 0, nil, nil, 0, 0}}},
		{"Example 2", args{n: 3}, [][]any{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
