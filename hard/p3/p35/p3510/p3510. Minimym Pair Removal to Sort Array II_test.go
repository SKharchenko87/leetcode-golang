package p3510

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPairRemoval(tt.args.nums); got != tt.want {
				t.Errorf("minimumPairRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
