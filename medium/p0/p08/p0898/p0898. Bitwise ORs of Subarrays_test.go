package p0898

import "testing"

func Test_subarrayBitwiseORs(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{arr: []int{0}}, 1},
		{"Example 2", args{arr: []int{1, 1, 2}}, 3},
		{"Example 3", args{arr: []int{1, 2, 4}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarrayBitwiseORs(tt.args.arr); got != tt.want {
				t.Errorf("subarrayBitwiseORs() = %v, want %v", got, tt.want)
			}
		})
	}
}

const N = 20000

var slice = make([]int, N)

func init() {
	for i := 0; i < N; i++ {
		slice[i] = 1 << (i % 10)
	}
}

func Benchmark_subarrayBitwiseORs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		subarrayBitwiseORs(slice)
	}
}

func Benchmark_subarrayBitwiseORs2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		subarrayBitwiseORs2(slice)
	}
}

func Benchmark_subarrayBitwiseORs1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		subarrayBitwiseORs1(slice)
	}
}

func Benchmark_subarrayBitwiseORs0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		subarrayBitwiseORs0(slice)
	}
}
