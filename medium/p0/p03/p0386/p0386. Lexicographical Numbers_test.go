package p0386

import (
	"reflect"
	"testing"
)

func Test_lexicalOrder(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{n: 13}, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"Example 2", args{n: 2}, []int{1, 2}},
		{"My 1", args{n: 23}, []int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2, 20, 21, 22, 23, 3, 4, 5, 6, 7, 8, 9}},
		{"My 2", args{n: 123}, []int{1, 10, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 11, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 12, 120, 121, 122, 123, 13, 14, 15, 16, 17, 18, 19, 2, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 3, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 4, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 5, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 6, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 7, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 8, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 9, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lexicalOrder(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{n: 13}, true},
		{"Example 1", args{n: 123}, true},
		{"Example 1", args{n: 323}, true},
		{"Example 1", args{n: 853}, true},
		{"Example 1", args{n: 487}, true},
		{"Example 1", args{n: 389}, true},
		{"Example 1", args{n: 1389}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.n); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lexicalOrder0(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"My 1", args{n: 123}, []int{1, 10, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 11, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 12, 120, 121, 122, 123, 13, 14, 15, 16, 17, 18, 19, 2, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 3, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 4, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 5, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 6, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 7, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 8, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 9, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lexicalOrder0(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicalOrder0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_lexicalOrder(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lexicalOrder(10000)
	}
}

func Benchmark_lexicalOrder0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lexicalOrder0(10000)
	}
}

func Benchmark_lexicalOrder1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lexicalOrder1(10000)
	}
}

func Benchmark_lexicalOrder2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lexicalOrder2(10000)
	}
}

func Benchmark_lexicalOrder3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lexicalOrder3(10000)
	}
}

func Benchmark_lexicalOrder4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lexicalOrder4(10000)
	}
}
