package p2785

import "testing"

func Test_sortVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s: "lEetcOde"}, "lEOtcede"},
		{"Example 2", args{s: "lYmpH"}, "lYmpH"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortVowels(tt.args.s); got != tt.want {
				t.Errorf("sortVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
