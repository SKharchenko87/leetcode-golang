package p0808

import (
	"testing"
)

func Test_soupServings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Example 1", args{n: 50}, 0.62500},
		{"Example 2", args{n: 100}, 0.71875},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soupServings(tt.args.n); got != tt.want {
				t.Errorf("soupServings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_soupServing(b *testing.B) {
	b.ReportAllocs()
	f := soupServings
	for i := 0; i < b.N; i++ {
		f(850)
		f(1000)
		f(4500)
	}
}

func Benchmark_soupServing3(b *testing.B) {
	b.ReportAllocs()
	f := soupServings3
	for i := 0; i < b.N; i++ {
		f(850)
		f(1000)
		f(4500)
	}
}

func Benchmark_soupServing2(b *testing.B) {
	b.ReportAllocs()
	f := soupServings2
	for i := 0; i < b.N; i++ {
		f(850)
		f(1000)
		f(4500)
	}
}

func Benchmark_soupServing1(b *testing.B) {
	b.ReportAllocs()
	f := soupServings1
	for i := 0; i < b.N; i++ {
		f(850)
		f(1000)
		f(4500)
	}
}

//TLE
//func Benchmark_soupServing0(b *testing.B) {
//	b.ReportAllocs()
//	f := soupServings0
//	for i := 0; i < b.N; i++ {
//		f(850)
//		f(1000)
//	}
//}
