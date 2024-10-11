package p0962

import "testing"

func Test_maxWidthRamp(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{6, 0, 8, 2, 1, 5}}, 4},
		{"Example 2", args{nums: []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}}, 7},
		{"Testcase 53", args{nums: []int{2, 2, 1}}, 1},
		{"Testcase 90", args{nums: []int{3, 4, 2, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWidthRamp(tt.args.nums); got != tt.want {
				t.Errorf("maxWidthRamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
