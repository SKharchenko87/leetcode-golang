package p2784

import "testing"

func Test_isGood(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{nums: []int{2, 1, 3}}, false},
		{"Example 2", args{nums: []int{1, 3, 3, 2}}, true},
		{"Example 3", args{nums: []int{1, 1}}, true},
		{"Example 4", args{nums: []int{3, 4, 4, 1, 2, 1}}, false},
		{"TestCase 790", args{nums: []int{60, 62, 16, 59, 8, 49, 51, 41, 42, 40, 46, 3, 10, 13, 53, 2, 63, 54, 32, 33, 31, 14, 12, 15, 1, 66, 61, 18, 52, 4, 55, 11, 26, 28, 47, 21, 25, 43, 65, 58, 45, 50, 17, 64, 23, 22, 30, 9, 38, 19, 24, 44, 37, 39, 48, 20, 35, 36, 27, 34, 56, 57, 29, 5, 7, 6, 66}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGood(tt.args.nums); got != tt.want {
				t.Errorf("isGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
