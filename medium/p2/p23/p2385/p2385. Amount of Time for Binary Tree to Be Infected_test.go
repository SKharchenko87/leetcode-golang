package p2385

import "testing"
import . "leetcode/stucture"

const null = NULL

func Test_amountOfTime(t *testing.T) {
	type args struct {
		root  *TreeNode
		start int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{SliceToTreeNodeFullMinInt([]int{1, 5, 3, null, 4, 10, 6, 9, 2}), 3}, 4},
		{"Case 2", args{SliceToTreeNodeFullMinInt([]int{1}), 1}, 0},
		{"Case 2", args{SliceToTreeNodeFullMinInt([]int{1, 2, null, 3, null, 4, null, 5}), 1}, 4},
		{"Case 2", args{SliceToTreeNodeFullMinInt([]int{1, null, 2, 3, 4, null, 5}), 4}, 3},
		{"Case 2", args{SliceToTreeNodeFullMinInt([]int{1, 2, null, 3, null, 4, null, 5}), 4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := amountOfTime(tt.args.root, tt.args.start); got != tt.want {
				t.Errorf("amountOfTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
