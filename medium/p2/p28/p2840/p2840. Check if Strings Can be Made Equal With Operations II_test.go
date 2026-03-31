package p2840

import "testing"

func Test_checkStrings(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{s1: "abcdba", s2: "cabdab"}, true},
		{"Example 2", args{s1: "abe", s2: "bea"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStrings(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(b *testing.B, f func(string, string) bool) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f("ababababababababababababababababababababababababababababababababababababababababababababababababababcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz",
			"bababababababababababababababababababababababababababababababababababababababababababababababababacbadfegihjklmonprqtsuvwxzyabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz",
		)
	}
}

func Benchmark_checkStrings(b *testing.B) {
	bench(b, checkStrings0)
}

func Benchmark_checkStrings0(b *testing.B) {
	bench(b, checkStrings0)
}
