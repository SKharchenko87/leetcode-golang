package p3867

import "testing"

func Test_gcdSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{2, 6, 4}}, 2},
		{"Example 2", args{nums: []int{3, 6, 2, 8}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdSum(tt.args.nums); got != tt.want {
				t.Errorf("gcdSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
