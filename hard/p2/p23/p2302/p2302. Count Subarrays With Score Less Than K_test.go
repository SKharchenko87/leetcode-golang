package p2302

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{2, 1, 4, 3, 5}, k: 10}, 6},
		{"Example 2", args{nums: []int{1, 1, 1}, k: 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
