package p1404

import (
	"testing"
)

func Test_numSteps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"1101"}, 6},
		{"Example 2", args{"10"}, 1},
		{"Example 3", args{"1"}, 0},
		{"My 1", args{"1000111"}, 11},
		{"My 2", args{"11"}, 3},
		{"My 3", args{"100"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSteps(tt.args.s); got != tt.want {
				t.Errorf("numSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(f func(s string) int, b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		f("11111100100010100010010110100111111110000010000010110110110101000001100100110111110001111001110010101001111101101011110011100110101101111001111111101101011110110000110010010011100100100101100100101110000101010101010101010000011101000100001100111101010001010110110011111000000000101010101010011111001111110010010001010000100101101000110000100011100110110011010100101100100110001011011000101001111110110110100111100111110110100111010001000001110111111001111000110111111100010001111000110101001101010000")
	}
}

func Benchmark_numSteps(b *testing.B) {
	bench(numSteps, b)
}

func Benchmark_numSteps5(b *testing.B) {
	bench(numSteps5, b)
}

func Benchmark_numSteps4(b *testing.B) {
	bench(numSteps4, b)
}

func Benchmark_numSteps3(b *testing.B) {
	bench(numSteps3, b)
}

func Benchmark_numSteps2(b *testing.B) {
	bench(numSteps2, b)
}

func Benchmark_numSteps1(b *testing.B) {
	bench(numSteps1, b)
}

func Benchmark_numSteps0(b *testing.B) {
	bench(numSteps0, b)
}
