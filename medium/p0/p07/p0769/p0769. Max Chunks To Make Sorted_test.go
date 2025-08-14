package p0769

import "testing"

func Test_maxChunksToSorted(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{arr: []int{4, 3, 2, 1, 0}}, 1},
		{"Example 2", args{arr: []int{1, 0, 2, 3, 4}}, 4},
		{"TestCase 36", args{arr: []int{2, 0, 1}}, 1},
		{"TestCase 50", args{arr: []int{1, 5, 3, 0, 2, 8, 7, 6, 4}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunksToSorted(tt.args.arr); got != tt.want {
				t.Errorf("maxChunksToSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
