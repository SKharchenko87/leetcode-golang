package p2486

import "testing"

func Test_appendCharacters(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "coaching", t: "coding"}, 4},
		{"Example 2", args{s: "abcde", t: "a"}, 0},
		{"Example 3", args{s: "z", t: "abcde"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendCharacters(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("appendCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
