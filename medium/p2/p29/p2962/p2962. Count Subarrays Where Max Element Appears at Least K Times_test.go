package p2962

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{1, 3, 2, 3, 3}, k: 2}, 6},
		{"Example 2", args{nums: []int{1, 4, 2, 1}, k: 3}, 0},
		{"TestCase 540", args{nums: []int{28, 5, 58, 91, 24, 91, 53, 9, 48, 85, 16, 70, 91, 91, 47, 91, 61, 4, 54, 61, 49}, k: 1}, 187},
		{"TestCase 603", args{nums: []int{73, 54, 15, 4, 23, 70, 53, 65, 73, 73, 2, 72, 36, 71, 73, 69, 35, 18, 62, 73, 62, 73, 73, 50, 30, 73, 20, 71, 60, 9, 12, 57, 48, 73, 40, 20, 8, 73, 73, 73, 34, 59, 31, 49, 73, 5, 51, 36, 47, 38, 36, 58, 34, 42, 23, 28, 52, 73}, k: 1}, 1537},
		{"My test 1", args{nums: []int{28, 5, 58, 91}, k: 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
