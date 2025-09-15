package p1935

import "testing"

func Test_canBeTypedWords(t *testing.T) {
	type args struct {
		text          string
		brokenLetters string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{text: "hello world", brokenLetters: "ad"}, 1},
		{"Example 2", args{text: "leet code", brokenLetters: "lt"}, 1},
		{"Example 3", args{text: "leet code", brokenLetters: "e"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeTypedWords(tt.args.text, tt.args.brokenLetters); got != tt.want {
				t.Errorf("canBeTypedWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
