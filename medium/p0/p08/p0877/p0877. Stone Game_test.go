package p0877

import "testing"

func Test_stoneGame(t *testing.T) {
	type args struct {
		piles []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{piles: []int{5, 3, 4, 5}}, true},
		{"Example 2", args{piles: []int{3, 7, 2, 3}}, true},
		{"TestCase 25", args{piles: []int{18, 3, 8, 1, 9, 7, 11, 13, 18, 11, 17, 20, 14, 2, 17, 20, 11, 14, 8, 7}}, true},
		{"TestCase 26", args{piles: []int{7, 7, 12, 16, 41, 48, 41, 48, 11, 9, 34, 2, 44, 30, 27, 12, 11, 39, 31, 8, 23, 11, 47, 25, 15, 23, 4, 17, 11, 50, 16, 50, 38, 34, 48, 27, 16, 24, 22, 48, 50, 10, 26, 27, 9, 43, 13, 42, 46, 24}}, true},
		{"Example 1", args{piles: []int{1, 10, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGame(tt.args.piles); got != tt.want {
				t.Errorf("stoneGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
