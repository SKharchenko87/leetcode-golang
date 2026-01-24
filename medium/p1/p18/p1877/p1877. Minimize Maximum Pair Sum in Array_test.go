package p1877

import "testing"

func Test_minPairSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{3, 5, 2, 3}}, 7},
		{"Example 2", args{nums: []int{3, 5, 4, 2, 4, 6}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPairSum(tt.args.nums); got != tt.want {
				t.Errorf("minPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(f func(nums []int) int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f([]int{9, 2, 10, 1, 10, 4, 8, 9, 7, 6, 8, 10, 8, 6, 5, 4, 3, 4, 2, 10})
	}
}

func Benchmark_minPairSum(b *testing.B) {
	bench(minPairSum, b)
}

func Benchmark_minPairSum0(b *testing.B) {
	bench(minPairSum0, b)
}
