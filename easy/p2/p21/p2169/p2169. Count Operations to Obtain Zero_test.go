package p2169

import "testing"

func Test_countOperations(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{num1: 2, num2: 3}, 3},
		{"Example 2", args{num1: 10, num2: 10}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOperations(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("countOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countOperations(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countOperations(100000, 1)
		countOperations(1, 100000)
	}
}

func Benchmark_countOperations2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countOperations2(100000, 1)
		countOperations2(1, 100000)
	}
}

func Benchmark_countOperations1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countOperations1(100000, 1)
		countOperations1(1, 100000)
	}
}

func Benchmark_countOperations0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		countOperations0(100000, 1)
		countOperations0(1, 100000)
	}
}
