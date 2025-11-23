package p1262

import "testing"

func Test_maxSumDivThree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{3, 6, 5, 1, 8}}, 18},
		{"Example 2", args{nums: []int{4}}, 0},
		{"Example 3", args{nums: []int{1, 2, 3, 4, 4}}, 12},
		{"TestCase 42", args{nums: []int{4, 1, 5, 3, 1}}, 12},
		{"Community 1", args{nums: []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 152, 11, 142, 2, 10, 12, 3}}, 927},
		{"Community 2", args{nums: []int{3, 7, 8, 9, 3, 1, 8, 5, 2, 10, 12, 3, 6, 5, 1, 7, 3, 15, 89, 144, 233}}, 573},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumDivThree(tt.args.nums); got != tt.want {
				t.Errorf("maxSumDivThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxSumDivThree(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxSumDivThree([]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 152, 11, 142, 2, 10, 12, 3})
		maxSumDivThree([]int{3, 7, 8, 9, 3, 1, 8, 5, 2, 10, 12, 3, 6, 5, 1, 7, 3, 15, 89, 144, 233})
	}
}

func Benchmark_maxSumDivThree4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxSumDivThree4([]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 152, 11, 142, 2, 10, 12, 3})
		maxSumDivThree4([]int{3, 7, 8, 9, 3, 1, 8, 5, 2, 10, 12, 3, 6, 5, 1, 7, 3, 15, 89, 144, 233})
	}
}

func Benchmark_maxSumDivThree3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxSumDivThree3([]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 152, 11, 142, 2, 10, 12, 3})
		maxSumDivThree3([]int{3, 7, 8, 9, 3, 1, 8, 5, 2, 10, 12, 3, 6, 5, 1, 7, 3, 15, 89, 144, 233})
	}
}

func Benchmark_maxSumDivThree2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxSumDivThree2([]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 152, 11, 142, 2, 10, 12, 3})
		maxSumDivThree2([]int{3, 7, 8, 9, 3, 1, 8, 5, 2, 10, 12, 3, 6, 5, 1, 7, 3, 15, 89, 144, 233})
	}
}

func Benchmark_maxSumDivThree1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxSumDivThree1([]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 152, 11, 142, 2, 10, 12, 3})
		maxSumDivThree1([]int{3, 7, 8, 9, 3, 1, 8, 5, 2, 10, 12, 3, 6, 5, 1, 7, 3, 15, 89, 144, 233})
	}
}

func Benchmark_maxSumDivThree0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxSumDivThree0([]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 152, 11, 142, 2, 10, 12, 3})
		maxSumDivThree0([]int{3, 7, 8, 9, 3, 1, 8, 5, 2, 10, 12, 3, 6, 5, 1, 7, 3, 15, 89, 144, 233})
	}
}

func Benchmark_Editorial1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxSumDivThreeEditorial1([]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 152, 11, 142, 2, 10, 12, 3})
		maxSumDivThreeEditorial1([]int{3, 7, 8, 9, 3, 1, 8, 5, 2, 10, 12, 3, 6, 5, 1, 7, 3, 15, 89, 144, 233})
	}
}

func Benchmark_Editorial2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxSumDivThreeEditorial2([]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 152, 11, 142, 2, 10, 12, 3})
		maxSumDivThreeEditorial2([]int{3, 7, 8, 9, 3, 1, 8, 5, 2, 10, 12, 3, 6, 5, 1, 7, 3, 15, 89, 144, 233})
	}
}

func Benchmark_Editorial3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxSumDivThreeEditorial3([]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 152, 11, 142, 2, 10, 12, 3})
		maxSumDivThreeEditorial3([]int{3, 7, 8, 9, 3, 1, 8, 5, 2, 10, 12, 3, 6, 5, 1, 7, 3, 15, 89, 144, 233})
	}
}
