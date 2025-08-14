package p0205

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{"egg", "add"}, true},
		{"Example 2", args{"foo", "bar"}, false},
		{"Example 3", args{"paper", "title"}, true},
		{"Example 24", args{"badc", "baba"}, false},
		{"Example 29", args{"13", "42"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
