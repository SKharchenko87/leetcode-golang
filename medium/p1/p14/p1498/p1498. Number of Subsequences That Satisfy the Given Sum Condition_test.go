package p1498

import "testing"

func Test_numSubseq(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{3, 5, 6, 7}, target: 9}, 4},
		{"Example 2", args{nums: []int{3, 3, 6, 8}, target: 10}, 6},
		{"Example 3", args{nums: []int{2, 3, 3, 4, 6, 7}, target: 12}, 61},
		{"TestCase 43", args{nums: []int{7, 10, 7, 5, 6, 7, 3, 4, 9, 6}, target: 9}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubseq(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("numSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
