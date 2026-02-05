package p3640

import "testing"

func Test_maxSumTrionic(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{0, -2, -1, -3, 0, 2, -1}}, -4},
		{"Example 2", args{nums: []int{1, 4, 2, 7}}, 14},
		{"My 1", args{nums: []int{0, -3, -2, -1, -3, 0, 2, -1}}, -4},
		{"My 2", args{nums: []int{0, -3, -2, -1, -3, 0, 2, 3, -1}}, -1},
		{"TestCase 331", args{nums: []int{467, 941, -696, -288}}, 424},
		{"TestCase 743", args{nums: []int{286, 528, -900, 327, 536, 625, 547, 997}}, 3032},
		{"TestCase 587", args{nums: []int{-432, 186, -568, 390}}, -424},
		{"TestCase 857", args{nums: []int{0, -2, -1, -3, 0, 0, 2, -1}}, -6},
		{"TestCase 735 ", args{nums: []int{1, 4, 2, 2, 3, 1, 2}}, 8},
		{"TestCase 735 ", args{nums: []int{-599, 720, -160, 799, 57, 261}}, 957},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumTrionic(tt.args.nums); got != tt.want {
				t.Errorf("maxSumTrionic() = %v, want %v", got, tt.want)
			}
		})
	}
}
