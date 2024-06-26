package p1248

import "testing"

func Test_numberOfSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 1, 2, 1, 1}, k: 3}, 2},
		{"Example 2", args{nums: []int{2, 4, 6}, k: 1}, 0},
		{"Example 3", args{nums: []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, k: 2}, 16},
		{"My 1", args{nums: []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2, 1}, k: 2}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numberOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
