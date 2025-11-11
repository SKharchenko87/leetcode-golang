package p0611

import "testing"

func Test_triangleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{2, 2, 3, 4}}, 3},
		{"Example 2", args{nums: []int{4, 2, 3, 4}}, 4},
		{"TestCase 55", args{nums: []int{1, 2, 3, 4, 5, 6}}, 7},
		{"TestCase 170", args{nums: []int{0, 1, 1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleNumber(tt.args.nums); got != tt.want {
				t.Errorf("trangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
