package p2392

import (
	"reflect"
	"testing"
)

func Test_buildMatrix(t *testing.T) {
	type args struct {
		k             int
		rowConditions [][]int
		colConditions [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{k: 3, rowConditions: [][]int{{1, 2}, {3, 2}}, colConditions: [][]int{{2, 1}, {3, 2}}}, [][]int{{3, 0, 0}, {0, 0, 1}, {0, 2, 0}}},
		{"Example 2", args{k: 3, rowConditions: [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}}, colConditions: [][]int{{2, 1}}}, [][]int{}},
		{"My 1", args{k: 3, rowConditions: [][]int{{1, 2}, {2, 3}, {3, 2}, {2, 3}}, colConditions: [][]int{{2, 1}}}, [][]int{}},
		{"Testcase 1", args{k: 8, rowConditions: [][]int{{1, 2}, {7, 3}, {4, 3}, {5, 8}, {7, 8}, {8, 2}, {5, 8}, {3, 2}, {1, 3}, {7, 6}, {4, 3}, {7, 4}, {4, 8}, {7, 3}, {7, 5}},
			colConditions: [][]int{{5, 7}, {2, 7}, {4, 3}, {6, 7}, {4, 3}, {2, 3}, {6, 2}}}, [][]int{{0, 0, 0, 0, 0, 0, 7, 0}, {0, 6, 0, 0, 0, 0, 0, 0}, {0, 0, 5, 0, 0, 0, 0, 0}, {0, 0, 0, 4, 0, 0, 0, 0}, {8, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 3, 0, 0}, {0, 0, 0, 0, 2, 0, 0, 0}}},
		{"Testcase 3", args{k: 6, rowConditions: [][]int{{1, 4}, {4, 3}, {1, 2}, {1, 4}}, colConditions: [][]int{{6, 2}, {5, 3}, {1, 3}, {5, 3}, {6, 4}, {2, 3}}}, [][]int{{0, 0, 0, 0, 0, 0, 7, 0}, {0, 6, 0, 0, 0, 0, 0, 0}, {0, 0, 5, 0, 0, 0, 0, 0}, {0, 0, 0, 4, 0, 0, 0, 0}, {8, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 3, 0, 0}, {0, 0, 0, 0, 2, 0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildMatrix(tt.args.k, tt.args.rowConditions, tt.args.colConditions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
