package p0719

import "testing"

func Test_smallestDistancePair(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 3, 1}, k: 1}, 0},
		{"Example 2", args{nums: []int{1, 1, 1}, k: 2}, 0},
		{"Example 3", args{nums: []int{1, 6, 1}, k: 3}, 5},
		{"TestCase 5", args{nums: []int{62, 100, 4}, k: 2}, 58},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestDistancePair(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("smallestDistancePair() = %v, want %v", got, tt.want)
			}
		})
	}
}
