package p0033

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, 4},
		{"Example 2", args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3}, -1},
		{"Example 3", args{nums: []int{1}, target: 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
