package p1255

import "testing"

func Test_maxScoreWords(t *testing.T) {
	type args struct {
		words   []string
		letters []byte
		score   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]string{"dog", "cat", "dad", "good"}, []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}, []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, 23},
		{"Example 2", args{[]string{"xxxz", "ax", "bx", "cx"}, []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}, []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}}, 27},
		{"Example 3", args{[]string{"leetcode"}, []byte{'l', 'e', 't', 'c', 'o', 'd'}, []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}}, 0},
		{"Example 38", args{[]string{"ac", "abd", "db", "ba", "dddd", "bca"}, []byte{'a', 'a', 'a', 'b', 'b', 'b', 'c', 'c', 'd', 'd', 'd', 'd'}, []int{6, 4, 4, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, 62},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScoreWords(tt.args.words, tt.args.letters, tt.args.score); got != tt.want {
				t.Errorf("maxScoreWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
