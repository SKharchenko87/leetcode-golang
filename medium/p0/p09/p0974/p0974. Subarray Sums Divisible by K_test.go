package p0974

import "testing"

func Test_subarraysDivByK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{4, 5, 0, -2, -3, 1}, 5}, 7},
		{"Example 2", args{[]int{5}, 9}, 0},
		{"Example ?", args{[]int{8, 9, 7, 8, 9}, 8}, 7},
		{"Example 46", args{[]int{-1, 2, 9}, 2}, 2},
		{"Example 46", args{[]int{-1, 3, 7}, 3}, 2},
		{"Example 52", args{[]int{7, -5, 5, -8, -6, 6, -4, 7, -8, -7}, 7}, 11},
		{"Example 46", args{[]int{-2, 4, -5, 7, 7}, 6}, 1},
		{"Example 40", args{[]int{4, 5, 0, -2, -3, 1}, 5}, 7},
		{"Example 60", args{[]int{204, -495, 612, 466, -946, 865, -928, -821, -415, -699, -528, -289, 471, -191, -461, -15, -215, -379, 801, 101, 955, 546, 141, -475, 374, -88, 443, -193, -8, -959, 480, -932, 362, 164, 552, -421, 372, -586, 59, -320, 957, 944, -582, 132, -341, 938, 849, 172, 440, 363, -293, 833, 799, -339, 495, -982, 633, 130, 905, -274, 252, -354, 566, 202, 975, -998, -719, -913, -749, 859, 898, -39, 894, 637, 558, 495, -625, 98, -152, 900, -653, -661, 318, -686, -780, 8, 594, 210, 510, -469, -430, -143, 982, 436, 893, 76, 109, 981, -620, 565}, 151}, 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
