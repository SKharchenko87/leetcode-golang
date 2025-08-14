package p3042

import "testing"

func Test_countPrefixSuffixPairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{words: []string{"a", "aba", "ababa", "aa"}}, 4},
		{"Example 2", args{words: []string{"pa", "papa", "ma", "mama"}}, 2},
		{"Example 3", args{words: []string{"abab", "ab"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrefixSuffixPairs4(tt.args.words); got != tt.want {
				t.Errorf("countPrefixSuffixPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countPrefixSuffixPairs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPrefixSuffixPairs([]string{"a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa"})
	}
}

func Benchmark_countPrefixSuffixPairs0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPrefixSuffixPairs0([]string{"a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa"})
	}
}

func Benchmark_countPrefixSuffixPairs1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPrefixSuffixPairs1([]string{"a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa"})
	}
}

func Benchmark_countPrefixSuffixPairs2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPrefixSuffixPairs2([]string{"a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa"})
	}
}

func Benchmark_countPrefixSuffixPairs3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPrefixSuffixPairs3([]string{"a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa"})
	}
}

func Benchmark_countPrefixSuffixPairs4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countPrefixSuffixPairs4([]string{"a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa"})
	}
}
