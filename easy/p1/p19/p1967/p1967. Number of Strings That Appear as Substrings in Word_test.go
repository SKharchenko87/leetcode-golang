package p1967

import "testing"

func Test_numOfStrings(t *testing.T) {
	type args struct {
		patterns []string
		word     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{patterns: []string{"a", "abc", "bc", "d"}, word: "abc"}, 3},
		{"Example 2", args{patterns: []string{"a", "b", "c"}, word: "aaaaabbbbb"}, 2},
		{"Example 3", args{patterns: []string{"a", "a", "a"}, word: "ab"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfStrings(tt.args.patterns, tt.args.word); got != tt.want {
				t.Errorf("numOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
