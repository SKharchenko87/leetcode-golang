package p2654

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{2, 6, 3, 4}}, 4},
		{"Example 2", args{nums: []int{2, 10, 6, 14}}, -1},
		{"My 1", args{nums: []int{2, 6, 9, 6, 2}}, 6},
		{"My 1", args{nums: []int{2, 42, 10, 12, 30, 15, 42}}, 9},
		{"Testcase 670", args{nums: []int{4, 2, 6, 3}}, 5},
		{"Testcase 741", args{nums: []int{1, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
