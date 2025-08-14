package p2560

import "testing"

func Test_minCapability(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{2, 3, 5, 9}, k: 2}, 5},
		{"Example 2", args{nums: []int{2, 7, 9, 3, 1}, k: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCapability(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minCapability() = %v, want %v", got, tt.want)
			}
		})
	}
}
