package p3163

import "testing"

func Test_compressedString(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{word: "abcde"}, "1a1b1c1d1e"},
		{"Example 2", args{word: "aaaaaaaaaaaaaabb"}, "9a5a2b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressedString(tt.args.word); got != tt.want {
				t.Errorf("compressedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
