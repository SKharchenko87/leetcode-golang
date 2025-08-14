package p3202

import "testing"

func Test_maximumLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 3, 4, 5}, k: 2}, 5},
		{"Example 2", args{nums: []int{1, 4, 2, 3, 1, 4}, k: 3}, 4},
		{"TestCase 391", args{nums: []int{1, 2, 3, 10, 2}, k: 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
