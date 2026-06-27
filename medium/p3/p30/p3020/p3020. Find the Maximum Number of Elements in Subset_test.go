package p3020

import "testing"

func Test_maximumLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{5, 4, 1, 2, 2}}, 3},
		{"Example 2", args{nums: []int{1, 3, 2, 4}}, 1},
		{"TestCase 2", args{nums: []int{1, 1}}, 1},
		{"TestCase 142", args{nums: []int{14, 14, 196, 196, 38416, 38416}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.nums); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
