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

func bench(f func(root *TreeNode) int, b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		f(SliceToTreeNodeFullMinInt([]int{1, 7, 0, 7, -8, NULL, NULL}))
		f(SliceToTreeNodeFullMinInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, -5, -6, -7, -8, 20, 30, 40, 50, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL}))
		f(SliceToTreeNodeFullMinInt([]int{100, -50, 200, -25, -25, 100, 100, -10, -15, -10, -15, 50, 50, 50, 50, -5, -5, -5, -5, -5, -5, -5, -5}))
	}
}

func Benchmark_maxLevelSum1(b *testing.B) {
	bench(maxLevelSum1, b)
}

func Benchmark_maxLevelSum0(b *testing.B) {
	bench(maxLevelSum0, b)
}

func Benchmark_maxLevelSum(b *testing.B) {
	bench(maxLevelSum, b)
}
