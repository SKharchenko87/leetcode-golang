package p2845

import "testing"

func Test_countInterestingSubarrays(t *testing.T) {
	type args struct {
		nums   []int
		modulo int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{3, 2, 4}, modulo: 2, k: 1}, 3},
		{"Example 2", args{nums: []int{3, 1, 9, 6}, modulo: 3, k: 0}, 2},
		{"TestCase 66", args{nums: []int{11, 12, 21, 31}, modulo: 10, k: 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countInterestingSubarrays(tt.args.nums, tt.args.modulo, tt.args.k); got != tt.want {
				t.Errorf("countInterestingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
