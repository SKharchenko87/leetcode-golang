package p2657

import (
	"reflect"
	"testing"
)

func Test_findThePrefixCommonArray(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{A: []int{1, 3, 2, 4}, B: []int{3, 1, 2, 4}}, []int{0, 2, 3, 4}},
		{"Example 2", args{A: []int{2, 3, 1}, B: []int{3, 1, 2}}, []int{0, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findThePrefixCommonArray(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findThePrefixCommonArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(b *testing.B, f func([]int, []int) []int) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f([]int{29, 47, 40, 19, 31, 34, 12, 18, 27, 37, 50, 13, 35, 16, 42, 22, 6, 32, 39, 5, 38, 15, 9, 25, 3, 24, 36, 23, 33, 11, 49, 21, 43, 45, 44, 4, 48, 28, 10, 26, 8, 46, 20, 1, 7, 14, 17, 41, 2, 30}, []int{13, 10, 44, 17, 3, 29, 16, 41, 45, 30, 25, 48, 18, 39, 42, 34, 2, 8, 6, 47, 36, 50, 31, 7, 21, 38, 19, 4, 32, 27, 26, 43, 46, 9, 40, 35, 12, 22, 20, 23, 11, 33, 5, 24, 28, 37, 1, 15, 49, 14})
		f([]int{24, 31, 37, 10, 38, 5, 32, 47, 42, 39, 6, 3, 33, 43, 21, 25, 18, 4, 9, 16, 13, 29, 23, 36, 35, 7, 48, 44, 49, 28, 34, 22, 46, 8, 20, 45, 14, 41, 12, 17, 2, 30, 15, 27, 1, 19, 40, 26, 11}, []int{46, 25, 23, 12, 45, 40, 6, 20, 35, 19, 33, 15, 39, 37, 27, 43, 47, 11, 13, 49, 41, 10, 18, 24, 4, 38, 44, 16, 2, 9, 1, 21, 34, 31, 3, 22, 48, 5, 7, 26, 32, 28, 36, 29, 42, 14, 30, 8, 17})
		f([]int{40, 11, 49, 3, 41, 32, 16, 23, 38, 35, 12, 28, 36, 21, 8, 33, 2, 14, 22, 5, 44, 37, 34, 27, 10, 47, 43, 26, 24, 48, 31, 19, 4, 42, 18, 30, 9, 39, 7, 45, 13, 20, 17, 25, 29, 15, 1, 6, 46}, []int{10, 9, 35, 47, 29, 11, 38, 4, 1, 14, 19, 39, 31, 44, 2, 8, 26, 18, 27, 23, 5, 28, 6, 13, 33, 30, 42, 12, 25, 15, 36, 17, 21, 34, 3, 43, 16, 49, 48, 24, 20, 37, 22, 32, 45, 7, 46, 40, 41})
	}
}

func Benchmark_findThePrefixCommonArray(b *testing.B) {
	bench(b, findThePrefixCommonArray)
}

func Benchmark_findThePrefixCommonArray0(b *testing.B) {
	bench(b, findThePrefixCommonArray0)
}

func Benchmark_findThePrefixCommonArray1(b *testing.B) {
	bench(b, findThePrefixCommonArray1)
}

func Benchmark_findThePrefixCommonArray2(b *testing.B) {
	bench(b, findThePrefixCommonArray2)
}
