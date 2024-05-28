package p1208

import "testing"

func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "abcd", t: "bcdf", maxCost: 3}, 3},
		{"Example 2", args{s: "abcd", t: "cdef", maxCost: 3}, 1},
		{"Example 3", args{s: "abcd", t: "acde", maxCost: 0}, 1},
		{"Example 8", args{s: "abcd", t: "cdef", maxCost: 1}, 0},
		{"Example 9", args{s: "pxezla", t: "loewbi", maxCost: 25}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
