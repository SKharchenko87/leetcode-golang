package p1460

import "testing"

func Test_canBeEqual(t *testing.T) {
	type args struct {
		target []int
		arr    []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{target: []int{1, 2, 3, 4}, arr: []int{2, 4, 1, 3}}, true},
		{"Example 2", args{target: []int{7}, arr: []int{7}}, true},
		{"Example 3", args{target: []int{3, 7, 9}, arr: []int{3, 7, 11}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeEqual(tt.args.target, tt.args.arr); got != tt.want {
				t.Errorf("canBeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
