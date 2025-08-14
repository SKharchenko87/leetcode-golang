package p2554

import "testing"

func Test_maxCount(t *testing.T) {
	type args struct {
		banned []int
		n      int
		maxSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{banned: []int{1, 6, 5}, n: 5, maxSum: 6}, 2},
		{"Example 2", args{banned: []int{1, 2, 3, 4, 5, 6, 7}, n: 8, maxSum: 1}, 0},
		{"Example 3", args{banned: []int{11}, n: 7, maxSum: 50}, 7},
		{"TestCase 105", args{banned: []int{87, 193, 85, 55, 14, 69, 26, 133, 171, 180, 4, 8, 29, 121, 182, 78, 157, 53, 26, 7, 117, 138, 57, 167, 8, 103, 32, 110, 15, 190, 139, 16, 49, 138, 68, 69, 92, 89, 140, 149, 107, 104, 2, 135, 193, 87, 21, 194, 192, 9, 161, 188, 73, 84, 83, 31, 86, 33, 138, 63, 127, 73, 114, 32, 66, 64, 19, 175, 108, 80, 176, 52, 124, 94, 33, 55, 130, 147, 39, 76, 22, 112, 113, 136, 100, 134, 155, 40, 170, 144, 37, 43, 151, 137, 82, 127, 73}, n: 1079, maxSum: 87}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCount(tt.args.banned, tt.args.n, tt.args.maxSum); got != tt.want {
				t.Errorf("maxCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
