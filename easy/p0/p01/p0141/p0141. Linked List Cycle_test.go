package p0141

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		head []int
		pos  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{[]int{3, 2, 0, -4}, 1}, true},
		{"Example 2", args{[]int{1, 2}, 0}, true},
		{"Example 3", args{[]int{1}, -1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.head, tt.args.pos); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
