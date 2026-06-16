package p3612

import "testing"

func Test_processStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s: "a#b%*"}, "ba"},
		{"Example 2", args{s: "z*#"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processStr(tt.args.s); got != tt.want {
				t.Errorf("processStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(b *testing.B, f func(s string) string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f("a###################")
	}
}

func Benchmark_processStr(b *testing.B) {
	bench(b, processStr)
}

func Benchmark_processStr0(b *testing.B) {
	bench(b, processStr0)
}
