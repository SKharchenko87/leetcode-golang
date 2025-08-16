package p1323

import "testing"

func Test_maximum69Number(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{num: 9669}, 9969},
		{"Example 2", args{num: 9996}, 9999},
		{"Example 3", args{num: 9999}, 9999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum69Number(tt.args.num); got != tt.want {
				t.Errorf("maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maximum69Number(b *testing.B) {
	b.ReportAllocs()
	f := maximum69Number
	for i := 0; i < b.N; i++ {
		_ = f(9669)
		_ = f(9996)
		_ = f(9999)
		_ = f(696)
	}
}

func Benchmark_maximum69Number4(b *testing.B) {
	b.ReportAllocs()
	f := maximum69Number4
	for i := 0; i < b.N; i++ {
		_ = f(9669)
		_ = f(9996)
		_ = f(9999)
		_ = f(696)
	}
}

func Benchmark_maximum69Number3(b *testing.B) {
	b.ReportAllocs()
	f := maximum69Number3
	for i := 0; i < b.N; i++ {
		_ = f(9669)
		_ = f(9996)
		_ = f(9999)
		_ = f(696)
	}
}

func Benchmark_maximum69Number2(b *testing.B) {
	b.ReportAllocs()
	f := maximum69Number2
	for i := 0; i < b.N; i++ {
		_ = f(9669)
		_ = f(9996)
		_ = f(9999)
		_ = f(696)
	}
}

func Benchmark_maximum69Number1(b *testing.B) {
	b.ReportAllocs()
	f := maximum69Number1
	for i := 0; i < b.N; i++ {
		_ = f(9669)
		_ = f(9996)
		_ = f(9999)
		_ = f(696)
	}
}

func Benchmark_maximum69Number0(b *testing.B) {
	b.ReportAllocs()
	f := maximum69Number0
	for i := 0; i < b.N; i++ {
		_ = f(9669)
		_ = f(9996)
		_ = f(9999)
		_ = f(696)
	}
}
