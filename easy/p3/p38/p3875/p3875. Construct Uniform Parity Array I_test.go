package p3875

import "testing"

func Test_uniformArray(t *testing.T) {
	type args struct {
		nums1 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{nums1: []int{2, 3}}, true},
		{"Example 2", args{nums1: []int{4, 6}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniformArray(tt.args.nums1); got != tt.want {
				t.Errorf("uniformArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
