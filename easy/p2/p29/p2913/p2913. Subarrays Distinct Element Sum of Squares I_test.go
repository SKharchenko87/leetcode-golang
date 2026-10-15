package p2913

import "testing"

func Test_sumCounts(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 1}}, 15},
		{"Example 2", args{nums: []int{1, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumCounts(tt.args.nums); got != tt.want {
				t.Errorf("sumCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
