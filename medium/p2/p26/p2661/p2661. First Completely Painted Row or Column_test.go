package p2661

import "testing"

func Test_firstCompleteIndex(t *testing.T) {
	type args struct {
		arr []int
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{arr: []int{1, 3, 4, 2}, mat: [][]int{{1, 4}, {2, 3}}}, 2},
		{"Example 2", args{arr: []int{2, 8, 7, 4, 1, 3, 5, 6, 9}, mat: [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstCompleteIndex(tt.args.arr, tt.args.mat); got != tt.want {
				t.Errorf("firstCompleteIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
