package p3120

import "testing"

func Test_numberOfSpecialChars(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{word: "aaAbcBC"}, 3},
		{"Example 2", args{word: "abc"}, 0},
		{"Example 3", args{word: "abBCab"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSpecialChars(tt.args.word); got != tt.want {
				t.Errorf("numberOfSpecialChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
