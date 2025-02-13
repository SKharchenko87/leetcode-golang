package p3066

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{2, 11, 10, 1, 3}, k: 10}, 2},
		{"Example 2", args{nums: []int{1, 1, 2, 4, 9}, k: 20}, 4},
		{"TestCase 2", args{nums: []int{999999999, 999999999, 999999999}, k: 1000000000}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
