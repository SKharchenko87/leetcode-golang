package p2779

import "testing"

func Test_maximumBeauty(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{4, 6, 1, 2}, k: 2}, 3},
		{"Example 2", args{nums: []int{1, 1, 1, 1}, k: 10}, 4},
		{"Example 3", args{nums: []int{38, 11, 31, 15, 50, 15}, k: 0}, 2},
		{"Example 3", args{nums: []int{49, 26}, k: 12}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBeauty(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}
