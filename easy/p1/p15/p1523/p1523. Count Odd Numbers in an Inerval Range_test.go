package p1523

import "testing"

func Test_countOdds(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{low: 3, high: 7}, 3},
		{"Example 2", args{low: 8, high: 10}, 1},
		{"TestCase 68", args{low: 14, high: 17}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOdds(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(f func(int, int) int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(0, 10000000)
	}
}

func Benchmark_countOdds(b *testing.B) {
	bench(countOdds, b)
}

func Benchmark_countOdds3(b *testing.B) {
	bench(countOdds3, b)
}
func Benchmark_countOdds2(b *testing.B) {
	bench(countOdds2, b)
}
func Benchmark_countOdds1(b *testing.B) {
	bench(countOdds1, b)
}
func Benchmark_countOdds0(b *testing.B) {
	bench(countOdds0, b)
}
