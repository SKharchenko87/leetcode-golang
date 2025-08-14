package p2373

import (
	"reflect"
	"testing"
)

func Test_largestLocal(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{[][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}}, [][]int{{9, 9}, {8, 6}}},
		{"Example 2", args{[][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}}, [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestLocal(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
