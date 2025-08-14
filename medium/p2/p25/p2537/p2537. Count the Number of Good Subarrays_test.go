package p2537

import "testing"

func Test_countGood(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{1, 1, 1, 1, 1}, k: 10}, 1},
		{"Example 2", args{nums: []int{3, 1, 4, 3, 2, 2, 4}, k: 2}, 4},
		{"My 1", args{nums: []int{1, 2, 2, 3, 1, 1, 1}, k: 3}, 6},
		{"My 1", args{nums: []int{3, 1, 9, 8, 7, 4, 3, 2, 4, 5, 6, 9, 8, 2, 4}, k: 2}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGood(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
