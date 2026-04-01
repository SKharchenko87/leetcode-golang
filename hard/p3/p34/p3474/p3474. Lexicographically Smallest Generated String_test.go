package p3474

import "testing"

func Test_generateString(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{str1: "TFTF", str2: "ab"}, "ababa"},
		{"Example 2", args{str1: "TFTF", str2: "abc"}, ""},
		{"Example 3", args{str1: "F", str2: "d"}, "a"},
		{"TestCase 698", args{str1: "FT", str2: "wvxyy"}, "awvxyy"},
		{"TestCase 723", args{str1: "TTFFT", str2: "fff"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateString(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("generateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
