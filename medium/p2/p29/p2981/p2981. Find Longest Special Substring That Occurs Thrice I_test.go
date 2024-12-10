package p2981

import "testing"

func Test_maximumLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "aaaa"}, 2},
		{"Example 2", args{s: "abcdef"}, -1},
		{"Example 3", args{s: "abcaba"}, 1},
		{"TestCase 759", args{s: "cbc"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.s); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
