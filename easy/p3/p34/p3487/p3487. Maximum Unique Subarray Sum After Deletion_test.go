package p3487

import "testing"

func Test_maxSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{1, 2, 3, 4, 5}}, 15},
		{"Example 2", args{nums: []int{1, 1, 0, 1, 1}}, 1},
		{"Example 3", args{nums: []int{1, 2, -1, -2, 1, 0, -1}}, 3},
		{"My 1", args{nums: []int{-1, -2, -1}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.nums); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxSum(b *testing.B) {
	b.ReportAllocs()
	var f func(nums []int) int
	f = maxSum
	for i := 0; i < b.N; i++ {
		f([]int{1, 2, 3, 4, 5})
		f([]int{1, 2, -1, -2, 1, 0, -1})
	}
}

func Benchmark_maxSum2(b *testing.B) {
	b.ReportAllocs()
	var f func(nums []int) int
	f = maxSum2
	for i := 0; i < b.N; i++ {
		f([]int{1, 2, 3, 4, 5})
		f([]int{1, 2, -1, -2, 1, 0, -1})
	}
}

func Benchmark_maxSum1(b *testing.B) {
	b.ReportAllocs()
	var f func(nums []int) int
	f = maxSum1
	for i := 0; i < b.N; i++ {
		f([]int{1, 2, 3, 4, 5})
		f([]int{1, 2, -1, -2, 1, 0, -1})
	}
}

func Benchmark_maxSum0(b *testing.B) {
	b.ReportAllocs()
	var f func(nums []int) int
	f = maxSum0
	for i := 0; i < b.N; i++ {
		f([]int{1, 2, 3, 4, 5})
		f([]int{1, 2, -1, -2, 1, 0, -1})
	}
}
