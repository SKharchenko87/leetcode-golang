package p3471

import "testing"

func Test_largestInteger(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{3, 9, 2, 1, 7}, k: 3}, 7},
		{"Example 2", args{nums: []int{3, 9, 7, 2, 1, 7}, k: 4}, 3},
		{"Example 3", args{nums: []int{0, 0}, k: 1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestInteger(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("largestInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
