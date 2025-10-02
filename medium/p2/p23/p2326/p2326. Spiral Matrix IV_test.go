package p2326

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		m    int
		n    int
		head []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{m: 3, n: 5, head: []int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0}}, [][]int{{3, 0, 2, 6, 8}, {5, 0, -1, -1, 1}, {5, 2, 4, 9, 7}}},
		{"Example 2", args{m: 1, n: 4, head: []int{0, 1, 2}}, [][]int{{0, 1, 2, -1}}},
		{"TestCase 10", args{m: 10, n: 8, head: []int{483, 100, 904, 632, 267, 352, 386, 887, 753}}, [][]int{{483, 100, 904, 632, 267, 352, 386, 887}, {-1, -1, -1, -1, -1, -1, -1, 753}, {-1, -1, -1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1, -1, -1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.m, tt.args.n, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
