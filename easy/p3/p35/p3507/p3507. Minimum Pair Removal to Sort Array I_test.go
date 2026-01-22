package p3507

import "testing"

func Test_minimumPairRemoval(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{5, 2, 3, 1}}, 2},
		{"Example 2", args{nums: []int{1, 2, 2}}, 0},
		{"TestCase 586", args{nums: []int{350, -113, -406, 764, -511, 90, -372, -411, 628, 822, -923, -146, 686, -631, -138, 157, -839, 302, 695, -436, 791, -920, -106, 802, 32, 483, 349, 346, 847, 704, -128, -495, 340, -316, -189, 585, -276}}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPairRemoval(tt.args.nums); got != tt.want {
				t.Errorf("minimumPairRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
