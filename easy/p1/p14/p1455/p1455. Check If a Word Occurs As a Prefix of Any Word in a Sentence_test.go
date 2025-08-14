package p1455

import "testing"

func Test_isPrefixOfWord(t *testing.T) {
	type args struct {
		sentence   string
		searchWord string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{sentence: "i love eating burger", searchWord: "burg"}, 4},
		{"Example 2", args{sentence: "this problem is an easy problem", searchWord: "pro"}, 2},
		{"Example 3", args{sentence: "i am tired", searchWord: "you"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrefixOfWord(tt.args.sentence, tt.args.searchWord); got != tt.want {
				t.Errorf("isPrefixOfWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
