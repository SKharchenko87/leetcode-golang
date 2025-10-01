package p2221

import "testing"

func Test_triangularSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 3, 4, 5}}, 8},
		{"Example 2", args{nums: []int{5}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangularSum(tt.args.nums); got != tt.want {
				t.Errorf("triangularSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
