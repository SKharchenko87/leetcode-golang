package p0085

import "testing"

func Test_maximalRectangle(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{matrix: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}}, 6},
		{"Example 2", args{matrix: [][]byte{{'0'}}}, 0},
		{"Example 3", args{matrix: [][]byte{{'1'}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(f func([][]byte) int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}})
	}
}

func Benchmark_maximalRectangle(b *testing.B) {
	bench(maximalRectangle, b)
}

func Benchmark_maximalRectangle0(b *testing.B) {
	bench(maximalRectangle0, b)
}
