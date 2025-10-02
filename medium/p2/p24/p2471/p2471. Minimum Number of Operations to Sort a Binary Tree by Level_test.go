package p2471

import "testing"
import . "leetcode/stucture"

func Test_run(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{root: []int{1, 4, 3, 7, 6, 8, 5, NULL, NULL, NULL, NULL, 9, NULL, 10}}, 3},
		{"Example 2", args{root: []int{1, 3, 2, 7, 6, 5, 4}}, 3},
		{"Example 3", args{root: []int{1, 2, 3, 4, 5, 6}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.root); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
