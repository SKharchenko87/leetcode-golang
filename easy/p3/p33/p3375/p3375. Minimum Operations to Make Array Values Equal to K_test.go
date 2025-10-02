package p3375

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
		{"Example 1", args{nums: []int{5, 2, 5, 4, 5}, k: 2}, 2},
		{"Example 2", args{nums: []int{2, 1, 2}, k: 2}, -1},
		{"Example 3", args{nums: []int{9, 7, 5, 3}, k: 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
