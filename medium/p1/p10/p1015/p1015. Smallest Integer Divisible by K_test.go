package p1015

import "testing"

func Test_smallestRepunitDivByK(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{k: 1}, 1},
		{"Example 2", args{k: 2}, -1},
		{"Example 3", args{k: 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRepunitDivByK(tt.args.k); got != tt.want {
				t.Errorf("smallestRepunitDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_smallestRepunitDivByK(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		smallestRepunitDivByK(23)
	}
}

func Benchmark_smallestRepunitDivByK2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		smallestRepunitDivByK2(23)
	}
}

func Benchmark_smallestRepunitDivByK1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		smallestRepunitDivByK1(23)
	}
}

func Benchmark_smallestRepunitDivByK0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		smallestRepunitDivByK0(23)
	}
}
