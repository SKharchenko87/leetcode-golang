package main

import (
	"reflect"
	"testing"
)

func Test_divideArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{nums: []int{1, 3, 4, 8, 7, 9, 3, 5, 1}, k: 2}, [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}}},
		{"Example 2", args{nums: []int{2, 4, 2, 2, 5, 2}, k: 2}, [][]int{}},
		{"Example 3", args{nums: []int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11}, k: 14}, [][]int{{2, 2, 12}, {4, 8, 5}, {5, 9, 7}, {7, 8, 5}, {5, 9, 10}, {11, 12, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideArray(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_divideArray(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		divideArray([]int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11}, 14)
	}
}

func Benchmark_divideArray2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		divideArray2([]int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11}, 14)
	}
}

func Benchmark_divideArray1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		divideArray1([]int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11}, 14)
	}
}

func Benchmark_divideArray0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		divideArray0([]int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11}, 14)
	}
}
