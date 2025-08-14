package p0725

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		head []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{head: []int{1, 2, 3}, k: 5}, [][]int{{1}, {2}, {3}, {}, {}}},
		{"Example 2", args{head: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, k: 3}, [][]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}}},
		{"TestCase 28", args{head: []int{0, 1, 2}, k: 2}, [][]int{{0, 1}, {2}}},
		{"TestCase 38", args{head: []int{0, 1, 2, 3, 4}, k: 3}, [][]int{{0, 1}, {2, 3}, {4}}},
		{"TestCase 38.2", args{head: []int{1, 2, 3, 4}, k: 5}, [][]int{{1}, {2}, {3}, {4}, {}}},
		{"TestCase 38.2", args{head: []int{}, k: 3}, [][]int{{}, {}, {}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
