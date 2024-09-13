package p1684

import "testing"

func Test_countConsistentStrings(t *testing.T) {
	type args struct {
		allowed string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{allowed: "ab", words: []string{"ad", "bd", "aaab", "baa", "badab"}}, 2},
		{"Example 2", args{allowed: "abc", words: []string{"a", "b", "c", "ab", "ac", "bc", "abc"}}, 7},
		{"Example 3", args{allowed: "cad", words: []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countConsistentStrings(tt.args.allowed, tt.args.words); got != tt.want {
				t.Errorf("countConsistentStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countConsistentStrings(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"})
	}
}

func Benchmark_countConsistentStrings0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countConsistentStrings0("ab", []string{"ad", "bd", "aaab", "baa", "badab"})
	}
}
