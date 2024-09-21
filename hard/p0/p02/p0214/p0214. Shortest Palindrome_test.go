package p0214

import "testing"

func Test_shortestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s: "aacecaaa"}, "aaacecaaa"},
		{"Example 2", args{s: "abcd"}, "dcbabcd"},
		{"TestCase 13", args{s: "aaaabbaa"}, "aabbaaaabbaa"},
		{"TestCase 27", args{s: "ba"}, "aba"},
		{"TestCase 27", args{s: "baa"}, "aabaa"},
		{"TestCase 114", args{s: "abbacd"}, "dcabbacd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("shortestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_stringRevertDeffer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stringRevertDeffer("baaabbaaaabbaasgjaaabbaaaabbaafgaaabbaaaabbaasdmgasaabbaaaabbaadmfasdjfasdasfgagasgd")
	}
}

func Benchmark_stringRevertSlice(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stringRevertSlice("baaabbaaaabbaasgjaaabbaaaabbaafgaaabbaaaabbaasdmgasaabbaaaabbaadmfasdjfasdasfgagasgd")
	}
}
