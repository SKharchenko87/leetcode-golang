package p2211

import "testing"

func Test_countCollisions(t *testing.T) {
	type args struct {
		directions string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{directions: "RLRSLL"}, 5},
		{"Example 2", args{directions: "LLRR"}, 0},
		{"TestCase 80", args{directions: "LLRLRLLSLRLLSLSSSS"}, 10},
		{"TestCase 80", args{directions: "L"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCollisions(tt.args.directions); got != tt.want {
				t.Errorf("countCollisions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countCollisions(b *testing.B) {
	b.ReportAllocs()
	f := countCollisions
	for i := 0; i < b.N; i++ {
		f("RSLRRLRSSLLRLLSRRLRRSLRSLRSSLRSLRLLRSLLRSSLRRLLLRRSLSLRRLLRSRLLRSRLSRLLRSLSRRLRRLSSRRLRL")
	}
}
func Benchmark_countCollisions5(b *testing.B) {
	b.ReportAllocs()
	f := countCollisions5
	for i := 0; i < b.N; i++ {
		f("RSLRRLRSSLLRLLSRRLRRSLRSLRSSLRSLRLLRSLLRSSLRRLLLRRSLSLRRLLRSRLLRSRLSRLLRSLSRRLRRLSSRRLRL")
	}
}
func Benchmark_countCollisions4(b *testing.B) {
	b.ReportAllocs()
	f := countCollisions4
	for i := 0; i < b.N; i++ {
		f("RSLRRLRSSLLRLLSRRLRRSLRSLRSSLRSLRLLRSLLRSSLRRLLLRRSLSLRRLLRSRLLRSRLSRLLRSLSRRLRRLSSRRLRL")
	}
}
func Benchmark_countCollisions3(b *testing.B) {
	b.ReportAllocs()
	f := countCollisions3
	for i := 0; i < b.N; i++ {
		f("RSLRRLRSSLLRLLSRRLRRSLRSLRSSLRSLRLLRSLLRSSLRRLLLRRSLSLRRLLRSRLLRSRLSRLLRSLSRRLRRLSSRRLRL")
	}
}
func Benchmark_countCollisions2(b *testing.B) {
	b.ReportAllocs()
	f := countCollisions2
	for i := 0; i < b.N; i++ {
		f("RSLRRLRSSLLRLLSRRLRRSLRSLRSSLRSLRLLRSLLRSSLRRLLLRRSLSLRRLLRSRLLRSRLSRLLRSLSRRLRRLSSRRLRL")
	}
}
func Benchmark_countCollisions1(b *testing.B) {
	b.ReportAllocs()
	f := countCollisions1
	for i := 0; i < b.N; i++ {
		f("RSLRRLRSSLLRLLSRRLRRSLRSLRSSLRSLRLLRSLLRSSLRRLLLRRSLSLRRLLRSRLLRSRLSRLLRSLSRRLRRLSSRRLRL")
	}
}
func Benchmark_countCollisions0(b *testing.B) {
	b.ReportAllocs()
	f := countCollisions0
	for i := 0; i < b.N; i++ {
		f("RSLRRLRSSLLRLLSRRLRRSLRSLRSSLRSLRLLRSLLRSSLRRLLLRRSLSLRRLLRSRLLRSRLSRLLRSLSRRLRRLSSRRLRL")
	}
}
