package p0885

import (
	"reflect"
	"testing"
)

func Test_spiralMatrixIII(t *testing.T) {
	type args struct {
		rows   int
		cols   int
		rStart int
		cStart int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{rows: 1, cols: 4, rStart: 0, cStart: 0}, [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
		{"Example 2", args{rows: 5, cols: 6, rStart: 1, cStart: 4}, [][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralMatrixIII(tt.args.rows, tt.args.cols, tt.args.rStart, tt.args.cStart); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralMatrixIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
