package p1061

import "testing"

func Test_smallestEquivalentString(t *testing.T) {
	type args struct {
		s1      string
		s2      string
		baseStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s1: "parker", s2: "morris", baseStr: "parser"}, "makkek"},
		{"Example 2", args{s1: "hello", s2: "world", baseStr: "hold"}, "hdld"},
		{"Example 3", args{s1: "leetcode", s2: "programs", baseStr: "sourcecode"}, "aauaaaaada"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestEquivalentString(tt.args.s1, tt.args.s2, tt.args.baseStr); got != tt.want {
				t.Errorf("smallestEquivalentString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_smallestEquivalentString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		smallestEquivalentString("leetcode", "programs", "sourcecode")
	}
}

func Benchmark_smallestEquivalentString1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		smallestEquivalentString1("leetcode", "programs", "sourcecode")
	}
}
