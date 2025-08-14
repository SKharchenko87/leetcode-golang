package p0073

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{"Example 2", args{matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{"TestCase 7", args{matrix: [][]int{{1}, {0}}}, [][]int{{0}, {0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
