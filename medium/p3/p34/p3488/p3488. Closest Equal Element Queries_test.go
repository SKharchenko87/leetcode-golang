package p3488

import (
	"reflect"
	"testing"
)

func Test_solveQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{nums: []int{1, 3, 1, 4, 1, 3, 2}, queries: []int{0, 3, 5}}, []int{2, -1, 3}},
		{"Example 2", args{nums: []int{1, 2, 3, 4}, queries: []int{0, 1, 2, 3}}, []int{-1, -1, -1, -1}},
		{"Testcase 421", args{nums: []int{6, 12, 17, 9, 16, 7, 6}, queries: []int{5, 6, 0, 4}}, []int{-1, 1, 1, -1}},
		{"Testcase 586", args{nums: []int{10}, queries: []int{0}}, []int{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveQueries(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(b *testing.B, f func([]int, []int) []int) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f([]int{9, 3, 9, 3, 3, 3, 11, 3, 18, 3, 7, 9, 11, 7, 6, 14, 3, 5}, []int{4, 1, 6, 3, 14, 13, 5, 12, 16, 9, 11})
	}
}

func BenchmarkSolveQueries(b *testing.B) {
	bench(b, solveQueries)
}

func BenchmarkSolveQueries1(b *testing.B) {
	bench(b, solveQueries1)
}

func BenchmarkSolveQueries0(b *testing.B) {
	bench(b, solveQueries0)
}
