package p2270

import "testing"

func Test_waysToSplitArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{10, 4, -8, 7}}, 2},
		{"Example 2", args{nums: []int{2, 3, 1, 0}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToSplitArray(tt.args.nums); got != tt.want {
				t.Errorf("waysToSplitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
