package p0788

import "testing"

func Test_rotatedDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 10}, 4},
		{"Example 2", args{n: 1}, 0},
		{"Example 3", args{n: 2}, 1},
		{"Example 3", args{n: 10000}, 2320},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotatedDigits(tt.args.n); got != tt.want {
				t.Errorf("rotatedDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(b *testing.B, f func(int) int) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(10000)
	}
}

func Benchmark_rotatedDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(b, rotatedDigits)
	}
}

func Benchmark_rotatedDigits0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(b, rotatedDigits0)
	}
}
