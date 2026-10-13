package p3727

import "testing"

func Test_maxAlternatingSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{1, 2, 3}}, 12},
		{"Example 2", args{nums: []int{1, -1, 2, -2, 3, -3}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAlternatingSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAlternatingSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
