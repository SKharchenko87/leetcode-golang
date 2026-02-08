package p0110

import "testing"
import . "leetcode/stucture"

func Test_run(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{root: []int{3, 9, 20, NULL, NULL, 15, 7}}, true},
		{"Example 2", args{root: []int{1, 2, 2, 3, 3, NULL, NULL, 4, 4}}, false},
		{"Example 3", args{root: []int{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.root); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
