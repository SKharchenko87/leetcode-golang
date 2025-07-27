package p2210

import "testing"

func Test_countHillValley(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{2, 4, 1, 1, 6, 5}}, 3},
		{"Example 2", args{nums: []int{6, 6, 5, 5, 4, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHillValley(tt.args.nums); got != tt.want {
				t.Errorf("countHillValley() = %v, want %v", got, tt.want)
			}
		})
	}
}
