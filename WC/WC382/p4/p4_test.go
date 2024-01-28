package p4

import (
	"testing"
)

func Test_minOrAfterOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{3, 5, 3, 2, 7}, 2}, 3},
		{"Case 2", args{[]int{7, 3, 15, 14, 2, 8}, 4}, 2},
		{"Case 3", args{[]int{10, 7, 10, 3, 9, 14, 9, 4}, 1}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOrAfterOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOrAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
