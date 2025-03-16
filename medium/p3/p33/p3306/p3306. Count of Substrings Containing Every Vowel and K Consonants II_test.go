package p3306

import "testing"

func Test_countOfSubstrings(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{word: "aeioqq", k: 1}, 0},
		{"Example 2", args{word: "aeiou", k: 0}, 1},
		{"Example 3", args{word: "ieaouqqieaouqq", k: 1}, 3},
		{"TestCase 626", args{word: "iqeaouqi", k: 2}, 3},
		{"TestCase 668", args{word: "aadieuoh", k: 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfSubstrings(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("countOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
