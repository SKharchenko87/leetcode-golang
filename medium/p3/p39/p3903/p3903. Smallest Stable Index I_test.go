package p3903

import "testing"

func Test_firstStableIndex(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{5, 0, 1, 4}, k: 3}, 3},
		{"Example 2", args{nums: []int{3, 2, 1}, k: 1}, -1},
		{"Example 3", args{nums: []int{0}, k: 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstStableIndex(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("firstStableIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
