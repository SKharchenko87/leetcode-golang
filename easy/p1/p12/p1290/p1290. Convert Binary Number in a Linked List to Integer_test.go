package p1290

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{head: []int{1, 0, 1}}, 5},
		{"Example 2", args{head: []int{0}}, 0},
		{"TestCase 12", args{head: []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}}, 18880},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.head); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
