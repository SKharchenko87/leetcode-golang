package p4

import "testing"

func Test_minimumCost(t *testing.T) {
	type args struct {
		nums []int
		k    int
		dist int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Case 1", args{[]int{1, 3, 2, 6, 4, 2}, 3, 3}, 5},
		{"Case 2", args{[]int{10, 1, 2, 2, 2, 1}, 4, 3}, 15},
		{"Case 3", args{[]int{10, 8, 18, 9}, 3, 1}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.nums, tt.args.k, tt.args.dist); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
