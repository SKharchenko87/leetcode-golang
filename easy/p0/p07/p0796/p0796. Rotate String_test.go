package p0796

import "testing"

func Test_rotateString(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{s: "abcde", goal: "cdeab"}, true},
		{"Example 2", args{s: "abcde", goal: "abced"}, false},
		{"TestCase 49", args{s: "defdefdefabcabc", goal: "defdefabcabcdef"}, true},
		{"TestCase 52", args{s: "bbbacddceeb", goal: "ceebbbbacdd"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(b *testing.B, f func(string, string) bool) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f("defdefdefabcabc", "defdefabcabcdef")
	}
}

func Benchmark_rotateString(b *testing.B) {
	bench(b, rotateString)
}

func Benchmark_rotateString1(b *testing.B) {
	bench(b, rotateString1)
}

func Benchmark_rotateString0(b *testing.B) {
	bench(b, rotateString0)
}
