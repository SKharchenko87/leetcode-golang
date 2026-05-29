package p1306

import "testing"

func Test_canReach(t *testing.T) {
	type args struct {
		arr   []int
		start int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{arr: []int{4, 2, 3, 0, 3, 1, 2}, start: 5}, true},
		{"Example 2", args{arr: []int{4, 2, 3, 0, 3, 1, 2}, start: 0}, true},
		{"Example 3", args{arr: []int{3, 0, 2, 1, 2}, start: 2}, false},
		{"TestCase 47", args{arr: []int{0, 1}, start: 1}, true},
		{"TestCase 48", args{arr: []int{1, 1, 1, 1, 1, 1, 1, 1, 0}, start: 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReach(tt.args.arr, tt.args.start); got != tt.want {
				t.Errorf("canReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
