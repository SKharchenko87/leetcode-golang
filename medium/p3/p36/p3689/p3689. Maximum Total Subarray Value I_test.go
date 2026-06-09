package p3689

import "testing"

func Test_maxTotalValue(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{1, 3, 2}, k: 2}, 4},
		{"Example 2", args{nums: []int{4, 2, 5, 1}, k: 3}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTotalValue(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxTotalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
