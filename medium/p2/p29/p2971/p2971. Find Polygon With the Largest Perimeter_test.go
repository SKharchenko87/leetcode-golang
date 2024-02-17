package p2971

import "testing"

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{[]int{5, 5, 5}}, 15},
		{"Example 2", args{[]int{1, 12, 1, 2, 5, 50, 3}}, 12},
		{"Example 3", args{[]int{5, 5, 50}}, -1},
		{"Example 1", args{[]int{5, 5, 11}}, -1},
		{"Example 1", args{[]int{1, 1, 2}}, -1},
		{"Example 1", args{[]int{1, 1, 5, 5}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.nums); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
