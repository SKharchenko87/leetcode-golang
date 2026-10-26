package p3942

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{0, 2, 1}}, 2},
		{"Example 2", args{nums: []int{1, 0, 2}}, 2},
		{"Example 3", args{nums: []int{2, 0, 1, 3}}, -1},
		{"TestCase 841", args{nums: []int{1, 2, 3, 0}}, 3},
		{"TestCase 998", args{nums: []int{0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
