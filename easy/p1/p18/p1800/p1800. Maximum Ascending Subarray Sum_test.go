package p1800

import "testing"

func Test_maxAscendingSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{10, 20, 30, 5, 10, 50}}, 65},
		{"Example 2", args{nums: []int{10, 20, 30, 40, 50}}, 150},
		{"Example 3", args{nums: []int{12, 17, 15, 13, 10, 11, 12}}, 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAscendingSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAscendingSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
