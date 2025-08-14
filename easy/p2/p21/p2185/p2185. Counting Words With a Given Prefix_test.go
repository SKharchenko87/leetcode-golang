package p2185

import "testing"

type args struct {
	words []string
	pref  string
}

func Test_prefixCount(t *testing.T) {

	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{words: []string{"pay", "attention", "practice", "attend"}, pref: "at"}, 2},
		{"Example 2", args{words: []string{"leetcode", "win", "loops", "success"}, pref: "code"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixCount(tt.args.words, tt.args.pref); got != tt.want {
				t.Errorf("prefixCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

var bench = args{words: []string{"pay", "attention", "practice", "attend"}, pref: "at"}

func Benchmark_prefixCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prefixCount([]string{"pay", "attention", "practice", "attend"}, "at")
	}
}

func Benchmark_prefixCount0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prefixCount0([]string{"pay", "attention", "practice", "attend"}, "at")
	}
}

func Benchmark_prefixCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prefixCount1([]string{"pay", "attention", "practice", "attend"}, "at")
	}
}

func Benchmark_prefixCount_(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prefixCount(bench.words, bench.pref)
	}
}

func Benchmark_prefixCount_0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prefixCount(bench.words, bench.pref)
	}
}

func Benchmark_prefixCount_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prefixCount(bench.words, bench.pref)
	}
}
