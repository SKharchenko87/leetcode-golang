package p0979

import "testing"
import . "leetcode/stucture"

func Test_distributeCoins(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{SliceToTreeNodeFullMinInt([]int{3, 0, 0})}, 2},
		{"Example 2", args{SliceToTreeNodeFullMinInt([]int{0, 3, 0})}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCoins(tt.args.root); got != tt.want {
				t.Errorf("distributeCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
