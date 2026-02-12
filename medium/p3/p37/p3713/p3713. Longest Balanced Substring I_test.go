package p3713

import "testing"

func Test_longestBalanced(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "abbac"}, 4},
		{"Example 2", args{s: "zzabccy"}, 4},
		{"Example 3", args{s: "aba"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestBalanced(tt.args.s); got != tt.want {
				t.Errorf("longestBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
