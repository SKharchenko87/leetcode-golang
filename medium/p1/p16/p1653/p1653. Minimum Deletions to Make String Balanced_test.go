package p1653

import "testing"

func Test_minimumDeletions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "aababbab"}, 2},
		{"Example 2", args{s: "bbaaaaabb"}, 2},
		{"Example 2", args{s: "bbaaaaabba"}, 3},
		{"Testcase 154", args{s: "bbbbbbbbbbbbbb"}, 0},
		{"Testcase 155", args{s: "bbbbbbbaabbbbbaaabbbabbbbaabbbbbbaabbaaabaabbbaaaabaaababbbabbabbaaaabbbabbbbbaabbababbbaaaaaababaaababaabbabbbaaaabbbbbabbabaaaabbbaba"}, 60},
		{"My 1", args{s: "bbabaaabbb"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions(tt.args.s); got != tt.want {
				t.Errorf("minimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(f func(string) int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f("bbbbbbbaabbbbbaaabbbabbbbaabbbbbbaabbaaabaabbbaaaabaaababbbabbabbaaaabbbabbbbbaabbababbbaaaaaababaaababaabbabbbaaaabbbbbabbabaaaabbbaba")
	}
}

func BenchmarkName(b *testing.B) {
	bench(minimumDeletions, b)
}
func BenchmarkName2(b *testing.B) {
	bench(minimumDeletions2, b)
}
func BenchmarkName1(b *testing.B) {
	bench(minimumDeletions1, b)
}
func BenchmarkName0(b *testing.B) {
	bench(minimumDeletions0, b)
}
