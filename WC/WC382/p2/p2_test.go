package p2

import (
	"testing"
)

func Test_maximumLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{5, 4, 1, 2, 2}}, 3},
		{"Case 2", args{[]int{1, 3, 2, 4}}, 1},
		{"Case my0", args{[]int{1, 3, 25, 5, 2, 25}}, 1},
		{"Case my1", args{[]int{1, 3, 5, 5, 2, 25}}, 3},
		{"Case my1", args{[]int{1, 3, 1, 1, 2, 25}}, 3},
		{"Case 3", args{[]int{1, 1}}, 1},
		{"Case 4", args{[]int{14, 14, 196, 196, 38416, 38416}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.nums); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
