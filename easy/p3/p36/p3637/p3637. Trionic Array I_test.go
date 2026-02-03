package p3637

import "testing"

func Test_isTrionic(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{nums: []int{1, 3, 5, 4, 2, 6}}, true},
		{"Example 2", args{nums: []int{2, 1, 3}}, false},
		{"TestCase 740", args{nums: []int{7, 6, 4, 4}}, false},
		{"TestCase 833", args{nums: []int{9, 4, 6, 8}}, false},
		{"TestCase 853", args{nums: []int{2, 4, 3, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTrionic(tt.args.nums); got != tt.want {
				t.Errorf("isTrionic() = %v, want %v", got, tt.want)
			}
		})
	}
}
