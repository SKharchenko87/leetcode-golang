package p2778

import "testing"

func Test_sumOfSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 3, 4}}, 21},
		{"Example 2", args{nums: []int{2, 7, 1, 19, 18, 3}}, 63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfSquares(tt.args.nums); got != tt.want {
				t.Errorf("sumOfSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
