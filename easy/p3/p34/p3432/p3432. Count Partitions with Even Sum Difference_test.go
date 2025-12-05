package p3432

import "testing"

func Test_countPartitions(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{10, 10, 3, 7, 6}}, 4},
		{"Example 2", args{nums: []int{1, 2, 2}}, 0},
		{"Example 3", args{nums: []int{2, 4, 6, 8}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPartitions(tt.args.nums); got != tt.want {
				t.Errorf("countPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

var benchArgs [][]int

func init() {
	l := 9
	benchArgs = make([][]int, l)
	x := 1
	for i := 0; i < l; i++ {
		for j := 1; j < x*10; j++ {

		}
		x *= 10
	}
}

func benchmark(f func([]int) int, b *testing.B) {
	b.ReportAllocs()
	args := make([][]int, len(benchArgs))
	for i := range benchArgs {
		args[i] = append([]int(nil), benchArgs[i]...)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, arg := range benchArgs {
			f(arg)
		}
	}
}

func Benchmark_countPartitions(b *testing.B)  { benchmark(countPartitions, b) }
func Benchmark_countPartitions2(b *testing.B) { benchmark(countPartitions2, b) }
func Benchmark_countPartitions1(b *testing.B) { benchmark(countPartitions1, b) }
func Benchmark_countPartitions0(b *testing.B) { benchmark(countPartitions0, b) }
