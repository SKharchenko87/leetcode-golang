package p0394

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Case 1", args{"3[a]2[bc]"}, "aaabcbc"},
		{"Case 2", args{"3[a2[c]]"}, "accaccacc"},
		{"Case 3", args{"2[abc]3[cd]ef"}, "abcabccdcdcdef"},
		{"Case 4", args{"3[z]2[2[y]pq4[2[jk]e1[f]]]ef"}, "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
