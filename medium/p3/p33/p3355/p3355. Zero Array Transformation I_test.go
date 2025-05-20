package p3355

import "testing"

func Test_isZeroArray(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{nums: []int{1, 0, 1}, queries: [][]int{{0, 2}}}, true},
		{"Example 2", args{nums: []int{4, 3, 2, 1}, queries: [][]int{{1, 3}, {0, 2}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZeroArray(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("isZeroArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
