package p1161

import "testing"
import . "leetcode/stucture"

func Test_maxLevelSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{SliceToTreeNodeFullMinInt([]int{1, 7, 0, 7, -8, NULL, NULL})},
			want: 2,
		},
		{
			name: "",
			args: args{SliceToTreeNodeFullMinInt([]int{989, NULL, 10250, 98693, -89388, NULL, NULL, NULL, -32127})},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLevelSum(tt.args.root); got != tt.want {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
