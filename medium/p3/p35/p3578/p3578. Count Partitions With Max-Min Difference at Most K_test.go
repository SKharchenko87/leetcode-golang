package p3578

import "testing"

func Test_countPartitions(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{9, 4, 1, 3, 7}, k: 4}, 6},
		{"Example 2", args{nums: []int{3, 3, 4}, k: 0}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPartitions(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
