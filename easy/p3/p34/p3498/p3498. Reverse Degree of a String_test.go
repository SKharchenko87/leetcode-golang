package p3498

import "testing"

func Test_reverseDegree(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "abc"}, 148},
		{"Example 2", args{s: "zaza"}, 160},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseDegree(tt.args.s); got != tt.want {
				t.Errorf("reverseDegree() = %v, want %v", got, tt.want)
			}
		})
	}
}
