package p2131

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{words: []string{"lc", "cl", "gg"}}, 6},
		{"Example 2", args{words: []string{"ab", "ty", "yt", "lc", "cl", "ab"}}, 8},
		{"Example 3", args{words: []string{"cc", "ll", "xx"}}, 2},
		{"TestCase 22", args{words: []string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}}, 22},
		{"Test 1", args{words: []string{"qo", "fo", "fq", "qf", "fo", "ff", "qq", "qf", "of", "of", "oo", "of", "of", "qf", "qf", "of"}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.words); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_longestPalindrome(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		longestPalindrome([]string{"qo", "fo", "fq", "qf", "fo", "ff", "qq", "qf", "of", "of", "oo", "of", "of", "qf", "qf", "of"})
	}
}

func Benchmark_longestPalindrome1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		longestPalindrome1([]string{"qo", "fo", "fq", "qf", "fo", "ff", "qq", "qf", "of", "of", "oo", "of", "of", "qf", "qf", "of"})
	}
}

func Benchmark_longestPalindrome0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		longestPalindrome0([]string{"qo", "fo", "fq", "qf", "fo", "ff", "qq", "qf", "of", "of", "oo", "of", "of", "qf", "qf", "of"})
	}
}
