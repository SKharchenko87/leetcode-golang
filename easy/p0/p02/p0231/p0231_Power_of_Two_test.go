package p0231

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{1}, true},
		{"Example 2", args{16}, true},
		{"Example 3", args{3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isPowerOfTwo(b *testing.B) {
	b.ReportAllocs()
	f := isPowerOfTwo
	for i := 0; i < b.N; i++ {
		f(2345333486)
		f(1024)
	}
}

func Benchmark_isPowerOfTwo5(b *testing.B) {
	b.ReportAllocs()
	f := isPowerOfTwo5
	for i := 0; i < b.N; i++ {
		f(2345333486)
		f(1024)
	}
}
func Benchmark_isPowerOfTwo4(b *testing.B) {
	b.ReportAllocs()
	f := isPowerOfTwo4
	for i := 0; i < b.N; i++ {
		f(2345333486)
		f(1024)
	}
}
func Benchmark_isPowerOfTwo3(b *testing.B) {
	b.ReportAllocs()
	f := isPowerOfTwo3
	for i := 0; i < b.N; i++ {
		f(2345333486)
		f(1024)
	}
}
func Benchmark_isPowerOfTwo2(b *testing.B) {
	b.ReportAllocs()
	f := isPowerOfTwo2
	for i := 0; i < b.N; i++ {
		f(2345333486)
		f(1024)
	}
}
func Benchmark_isPowerOfTwo1(b *testing.B) {
	b.ReportAllocs()
	f := isPowerOfTwo1
	for i := 0; i < b.N; i++ {
		f(2345333486)
		f(1024)
	}
}
