package p3445

import "testing"

func Test_maxDifference(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "12233", k: 4}, -1},
		{"Example 2", args{s: "1122211", k: 3}, 1},
		{"Example 3", args{s: "110", k: 3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDifference(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
