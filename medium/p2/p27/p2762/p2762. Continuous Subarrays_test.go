package p2762

import "testing"

func Test_continuousSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{nums: []int{5, 4, 2, 4}}, 8},
		{"Example 2", args{nums: []int{1, 2, 3}}, 6},
		{"Example 1918", args{nums: []int{67, 67, 66, 67, 68, 69, 70, 71, 71}}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := continuousSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("continuousSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
