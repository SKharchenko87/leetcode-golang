package p2

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
		// TODO: Add test cases.
		{"Case 1", args{"aaaa"}, 2},
		{"Case 1", args{"abcdef"}, -1},
		{"Case 1", args{"abcaba"}, 1},
		{"Case 1", args{"fafff"}, 1},
		{"Case 1", args{"ceeeeeeeeeeeebmmmfffeeeeeeeeeeeewww"}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.s); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
