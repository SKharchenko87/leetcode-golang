package p1022

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{root: []int{1, 0, 1, 0, 1, 0, 1}}, 22},
		{"Example 2", args{root: []int{0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.root); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
