package p3425

import (
	"reflect"
	"testing"
)

func Test_longestSpecialPath(t *testing.T) {
	type args struct {
		edges [][]int
		nums  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{edges: [][]int{{0, 1, 2}, {1, 2, 3}, {1, 3, 5}, {1, 4, 4}, {2, 5, 6}}, nums: []int{2, 1, 2, 1, 3, 1}}, []int{6, 2}},
		{"Example 2", args{edges: [][]int{{1, 0, 8}}, nums: []int{2, 2}}, []int{0, 1}},
		{"Example 30", args{edges: [][]int{{1, 0, 7}, {1, 2, 4}}, nums: []int{1, 1, 3}}, []int{4, 2}},
		{"Example 56", args{edges: [][]int{{1, 0, 5}, {0, 2, 10}}, nums: []int{3, 3, 5}}, []int{10, 2}},
		{"Example 90", args{edges: [][]int{{0, 1, 10}, {1, 2, 5}}, nums: []int{2, 1, 5}}, []int{15, 3}},
		{"Example 52", args{edges: [][]int{{1, 0, 5}, {2, 1, 3}}, nums: []int{3, 1, 1}}, []int{5, 2}},
		{"Example 169", args{edges: [][]int{{1, 0, 1}, {1, 2, 6}}, nums: []int{4, 3, 5}}, []int{7, 3}},
		{"Example 169", args{edges: [][]int{{2, 0, 2}, {1, 2, 10}}, nums: []int{1, 5, 4}}, []int{12, 3}},
		{"Example 498", args{edges: [][]int{{3, 0, 5}, {3, 1, 6}, {2, 3, 5}}, nums: []int{2, 2, 2, 1}}, []int{6, 2}},
		{"Example 554", args{edges: [][]int{{1, 0, 9}, {2, 1, 7}, {3, 2, 10}}, nums: []int{3, 9, 18, 18}}, []int{16, 3}},
		{"Example 565", args{edges: [][]int{{1, 2, 4}, {0, 1, 8}, {0, 3, 6}}, nums: []int{1, 3, 1, 1}}, []int{8, 2}},
		{"Example 637", args{edges: [][]int{{1, 0, 9}, {1, 2, 6}, {2, 3, 4}}, nums: []int{1, 4, 4, 3}}, []int{9, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSpecialPath(tt.args.edges, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestSpecialPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
