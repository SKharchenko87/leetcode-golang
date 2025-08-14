package p0802

import (
	"reflect"
	"testing"
)

func Test_eventualSafeNodes(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{graph: [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}}, []int{2, 4, 5, 6}},
		{"Example 2", args{graph: [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}}, []int{4}},
		{"TestCase 84", args{graph: [][]int{{2, 3, 6, 8, 13}, {7}, {7}, {12}, {9, 19}, {}, {7, 12, 13, 15, 17}, {5, 12, 19}, {12, 13, 16, 17}, {11, 14, 15, 16, 17, 19}, {10, 11, 16, 18}, {}, {}, {15, 16}, {1, 15, 16, 17, 18, 19}, {16, 17, 18, 19}, {17, 18, 19}, {18}, {19}, {}}}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eventualSafeNodes(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("eventualSafeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
