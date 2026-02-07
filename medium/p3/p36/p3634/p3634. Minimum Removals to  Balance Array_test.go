package p3634

import "testing"

func Test_minRemoval(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{2, 1, 5}, k: 2}, 1},
		{"Example 2", args{nums: []int{1, 6, 2, 9}, k: 3}, 2},
		{"Example 3", args{nums: []int{4, 6}, k: 2}, 0},
		{"TestCase 427", args{nums: []int{1, 34, 23}, k: 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoval(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
