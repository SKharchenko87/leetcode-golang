package p3151

import "testing"

func Test_isArraySpecial(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{nums: []int{1}}, true},
		{"Example 2", args{nums: []int{2, 1, 4}}, true},
		{"Example 3", args{nums: []int{4, 3, 1, 6}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial(tt.args.nums); got != tt.want {
				t.Errorf("isArraySpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
