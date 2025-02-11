package p1910

import "testing"

func Test_removeOccurrences(t *testing.T) {
	type args struct {
		s    string
		part string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s: "daabcbaabcbc", part: "abc"}, "dab"},
		{"Example 2", args{s: "axxxxyyyyb", part: "xy"}, "ab"},
		{"TestCase 2", args{s: "eemckxmckx", part: "emckx"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOccurrences(tt.args.s, tt.args.part); got != tt.want {
				t.Errorf("removeOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
