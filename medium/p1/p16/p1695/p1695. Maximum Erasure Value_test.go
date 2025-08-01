package p1695

import "testing"

func Test_maximumUniqueSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{4, 2, 4, 5, 6}}, 17},
		{"Example 2", args{nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5}}, 8},
		{"TestCase 57", args{nums: []int{10000}}, 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUniqueSubarray(tt.args.nums); got != tt.want {
				t.Errorf("maximumUniqueSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maximumUniqueSubarray(b *testing.B) {
	b.ReportAllocs()
	f := func(ints []int) int {
		return maximumUniqueSubarray(ints)
	}
	for i := 0; i < b.N; i++ {
		f([]int{8, 15, 6, 12, 3, 4, 2, 15, 10, 20, 7, 14, 5, 17, 1, 20, 7, 2, 3, 16, 13, 19, 8, 15, 16, 15, 7, 15, 18, 9, 17, 3, 8, 14, 4, 9, 13, 19, 1, 7, 8, 15, 19, 2, 10, 20, 7, 16, 6, 8})
		f([]int{8, 15, 6, 12, 3, 4, 2, 15, 10, 20, 7, 14, 5, 17, 1, 20, 7, 2, 3, 16, 13, 19, 8, 15, 16, 15, 7, 15, 18, 9, 17, 3, 8, 14, 4, 9, 13, 19, 1, 7, 8, 15, 19, 2, 10, 20, 7, 16, 6, 8})
		f([]int{1, 10, 3, 8, 19, 5, 12, 3, 17, 3, 8, 18, 19, 9, 5, 11, 1, 6, 8, 1, 7, 20, 2, 8, 11, 15, 4, 14, 11, 13, 20, 18, 6, 16, 5, 11, 8, 1, 5, 13, 19, 11, 15, 9, 8, 4, 4, 6, 7, 2})
		f([]int{10, 17, 10, 7, 4, 12, 16, 4, 20, 13, 7, 18, 9, 16, 20, 3, 5, 6, 1, 6, 12, 16, 15, 10, 13, 3, 4, 17, 1, 11, 11, 6, 15, 12, 1, 17, 11, 10, 2, 18, 15, 4, 5, 1, 14, 15, 9, 8, 8, 5})
		f([]int{10, 1, 4, 15, 14, 14, 1, 6, 11, 14, 1, 9, 12, 17, 4, 7, 8, 1, 1, 5, 20, 9, 6, 17, 17, 11, 5, 18, 7, 17, 13, 10, 17, 20, 5, 10, 9, 3, 17, 2, 1, 2, 20, 11, 3, 10, 7, 15, 1, 19})
		f([]int{18, 8, 8, 19, 3, 1, 19, 1, 15, 13, 10, 4, 12, 2, 3, 3, 11, 4, 20, 11, 12, 4, 14, 9, 8, 1, 13, 15, 9, 20, 12, 13, 11, 8, 7, 8, 15, 10, 5, 17, 4, 10, 18, 5, 12, 18, 5, 16, 13, 2})
	}
}

func Benchmark_maximumUniqueSubarray3(b *testing.B) {
	b.ReportAllocs()
	f := func(ints []int) int {
		return maximumUniqueSubarray3(ints)
	}
	for i := 0; i < b.N; i++ {
		f([]int{8, 15, 6, 12, 3, 4, 2, 15, 10, 20, 7, 14, 5, 17, 1, 20, 7, 2, 3, 16, 13, 19, 8, 15, 16, 15, 7, 15, 18, 9, 17, 3, 8, 14, 4, 9, 13, 19, 1, 7, 8, 15, 19, 2, 10, 20, 7, 16, 6, 8})
		f([]int{8, 15, 6, 12, 3, 4, 2, 15, 10, 20, 7, 14, 5, 17, 1, 20, 7, 2, 3, 16, 13, 19, 8, 15, 16, 15, 7, 15, 18, 9, 17, 3, 8, 14, 4, 9, 13, 19, 1, 7, 8, 15, 19, 2, 10, 20, 7, 16, 6, 8})
		f([]int{1, 10, 3, 8, 19, 5, 12, 3, 17, 3, 8, 18, 19, 9, 5, 11, 1, 6, 8, 1, 7, 20, 2, 8, 11, 15, 4, 14, 11, 13, 20, 18, 6, 16, 5, 11, 8, 1, 5, 13, 19, 11, 15, 9, 8, 4, 4, 6, 7, 2})
		f([]int{10, 17, 10, 7, 4, 12, 16, 4, 20, 13, 7, 18, 9, 16, 20, 3, 5, 6, 1, 6, 12, 16, 15, 10, 13, 3, 4, 17, 1, 11, 11, 6, 15, 12, 1, 17, 11, 10, 2, 18, 15, 4, 5, 1, 14, 15, 9, 8, 8, 5})
		f([]int{10, 1, 4, 15, 14, 14, 1, 6, 11, 14, 1, 9, 12, 17, 4, 7, 8, 1, 1, 5, 20, 9, 6, 17, 17, 11, 5, 18, 7, 17, 13, 10, 17, 20, 5, 10, 9, 3, 17, 2, 1, 2, 20, 11, 3, 10, 7, 15, 1, 19})
		f([]int{18, 8, 8, 19, 3, 1, 19, 1, 15, 13, 10, 4, 12, 2, 3, 3, 11, 4, 20, 11, 12, 4, 14, 9, 8, 1, 13, 15, 9, 20, 12, 13, 11, 8, 7, 8, 15, 10, 5, 17, 4, 10, 18, 5, 12, 18, 5, 16, 13, 2})
	}
}

func Benchmark_maximumUniqueSubarray2(b *testing.B) {
	b.ReportAllocs()
	f := func(ints []int) int {
		return maximumUniqueSubarray2(ints)
	}
	for i := 0; i < b.N; i++ {
		f([]int{8, 15, 6, 12, 3, 4, 2, 15, 10, 20, 7, 14, 5, 17, 1, 20, 7, 2, 3, 16, 13, 19, 8, 15, 16, 15, 7, 15, 18, 9, 17, 3, 8, 14, 4, 9, 13, 19, 1, 7, 8, 15, 19, 2, 10, 20, 7, 16, 6, 8})
		f([]int{8, 15, 6, 12, 3, 4, 2, 15, 10, 20, 7, 14, 5, 17, 1, 20, 7, 2, 3, 16, 13, 19, 8, 15, 16, 15, 7, 15, 18, 9, 17, 3, 8, 14, 4, 9, 13, 19, 1, 7, 8, 15, 19, 2, 10, 20, 7, 16, 6, 8})
		f([]int{1, 10, 3, 8, 19, 5, 12, 3, 17, 3, 8, 18, 19, 9, 5, 11, 1, 6, 8, 1, 7, 20, 2, 8, 11, 15, 4, 14, 11, 13, 20, 18, 6, 16, 5, 11, 8, 1, 5, 13, 19, 11, 15, 9, 8, 4, 4, 6, 7, 2})
		f([]int{10, 17, 10, 7, 4, 12, 16, 4, 20, 13, 7, 18, 9, 16, 20, 3, 5, 6, 1, 6, 12, 16, 15, 10, 13, 3, 4, 17, 1, 11, 11, 6, 15, 12, 1, 17, 11, 10, 2, 18, 15, 4, 5, 1, 14, 15, 9, 8, 8, 5})
		f([]int{10, 1, 4, 15, 14, 14, 1, 6, 11, 14, 1, 9, 12, 17, 4, 7, 8, 1, 1, 5, 20, 9, 6, 17, 17, 11, 5, 18, 7, 17, 13, 10, 17, 20, 5, 10, 9, 3, 17, 2, 1, 2, 20, 11, 3, 10, 7, 15, 1, 19})
		f([]int{18, 8, 8, 19, 3, 1, 19, 1, 15, 13, 10, 4, 12, 2, 3, 3, 11, 4, 20, 11, 12, 4, 14, 9, 8, 1, 13, 15, 9, 20, 12, 13, 11, 8, 7, 8, 15, 10, 5, 17, 4, 10, 18, 5, 12, 18, 5, 16, 13, 2})
	}
}

func Benchmark_maximumUniqueSubarray1(b *testing.B) {
	b.ReportAllocs()
	f := func(ints []int) int {
		return maximumUniqueSubarray1(ints)
	}
	for i := 0; i < b.N; i++ {
		f([]int{8, 15, 6, 12, 3, 4, 2, 15, 10, 20, 7, 14, 5, 17, 1, 20, 7, 2, 3, 16, 13, 19, 8, 15, 16, 15, 7, 15, 18, 9, 17, 3, 8, 14, 4, 9, 13, 19, 1, 7, 8, 15, 19, 2, 10, 20, 7, 16, 6, 8})
		f([]int{8, 15, 6, 12, 3, 4, 2, 15, 10, 20, 7, 14, 5, 17, 1, 20, 7, 2, 3, 16, 13, 19, 8, 15, 16, 15, 7, 15, 18, 9, 17, 3, 8, 14, 4, 9, 13, 19, 1, 7, 8, 15, 19, 2, 10, 20, 7, 16, 6, 8})
		f([]int{1, 10, 3, 8, 19, 5, 12, 3, 17, 3, 8, 18, 19, 9, 5, 11, 1, 6, 8, 1, 7, 20, 2, 8, 11, 15, 4, 14, 11, 13, 20, 18, 6, 16, 5, 11, 8, 1, 5, 13, 19, 11, 15, 9, 8, 4, 4, 6, 7, 2})
		f([]int{10, 17, 10, 7, 4, 12, 16, 4, 20, 13, 7, 18, 9, 16, 20, 3, 5, 6, 1, 6, 12, 16, 15, 10, 13, 3, 4, 17, 1, 11, 11, 6, 15, 12, 1, 17, 11, 10, 2, 18, 15, 4, 5, 1, 14, 15, 9, 8, 8, 5})
		f([]int{10, 1, 4, 15, 14, 14, 1, 6, 11, 14, 1, 9, 12, 17, 4, 7, 8, 1, 1, 5, 20, 9, 6, 17, 17, 11, 5, 18, 7, 17, 13, 10, 17, 20, 5, 10, 9, 3, 17, 2, 1, 2, 20, 11, 3, 10, 7, 15, 1, 19})
		f([]int{18, 8, 8, 19, 3, 1, 19, 1, 15, 13, 10, 4, 12, 2, 3, 3, 11, 4, 20, 11, 12, 4, 14, 9, 8, 1, 13, 15, 9, 20, 12, 13, 11, 8, 7, 8, 15, 10, 5, 17, 4, 10, 18, 5, 12, 18, 5, 16, 13, 2})
	}
}

func Benchmark_maximumUniqueSubarray0(b *testing.B) {
	b.ReportAllocs()
	f := func(ints []int) int {
		return maximumUniqueSubarray0(ints)
	}
	for i := 0; i < b.N; i++ {
		f([]int{8, 15, 6, 12, 3, 4, 2, 15, 10, 20, 7, 14, 5, 17, 1, 20, 7, 2, 3, 16, 13, 19, 8, 15, 16, 15, 7, 15, 18, 9, 17, 3, 8, 14, 4, 9, 13, 19, 1, 7, 8, 15, 19, 2, 10, 20, 7, 16, 6, 8})
		f([]int{8, 15, 6, 12, 3, 4, 2, 15, 10, 20, 7, 14, 5, 17, 1, 20, 7, 2, 3, 16, 13, 19, 8, 15, 16, 15, 7, 15, 18, 9, 17, 3, 8, 14, 4, 9, 13, 19, 1, 7, 8, 15, 19, 2, 10, 20, 7, 16, 6, 8})
		f([]int{1, 10, 3, 8, 19, 5, 12, 3, 17, 3, 8, 18, 19, 9, 5, 11, 1, 6, 8, 1, 7, 20, 2, 8, 11, 15, 4, 14, 11, 13, 20, 18, 6, 16, 5, 11, 8, 1, 5, 13, 19, 11, 15, 9, 8, 4, 4, 6, 7, 2})
		f([]int{10, 17, 10, 7, 4, 12, 16, 4, 20, 13, 7, 18, 9, 16, 20, 3, 5, 6, 1, 6, 12, 16, 15, 10, 13, 3, 4, 17, 1, 11, 11, 6, 15, 12, 1, 17, 11, 10, 2, 18, 15, 4, 5, 1, 14, 15, 9, 8, 8, 5})
		f([]int{10, 1, 4, 15, 14, 14, 1, 6, 11, 14, 1, 9, 12, 17, 4, 7, 8, 1, 1, 5, 20, 9, 6, 17, 17, 11, 5, 18, 7, 17, 13, 10, 17, 20, 5, 10, 9, 3, 17, 2, 1, 2, 20, 11, 3, 10, 7, 15, 1, 19})
		f([]int{18, 8, 8, 19, 3, 1, 19, 1, 15, 13, 10, 4, 12, 2, 3, 3, 11, 4, 20, 11, 12, 4, 14, 9, 8, 1, 13, 15, 9, 20, 12, 13, 11, 8, 7, 8, 15, 10, 5, 17, 4, 10, 18, 5, 12, 18, 5, 16, 13, 2})
	}
}
