package p0070

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{2}, 2},
		{"Case 2", args{3}, 3},
		{"Case 2", args{4}, 5},
		{"My 45", args{45}, 1836311903},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs0(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{2}, 2},
		{"Case 2", args{3}, 3},
		{"Case 2", args{4}, 5},
		{"My 45", args{45}, 1836311903},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs0(tt.args.n); got != tt.want {
				t.Errorf("climbStairs0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{2}, 2},
		{"Case 2", args{3}, 3},
		{"Case 2", args{4}, 5},
		{"My 45", args{45}, 1836311903},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs1(tt.args.n); got != tt.want {
				t.Errorf("climbStairs1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_climbStairs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		climbStairs(45)
	}
}

func Benchmark_climbStairs0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		climbStairs0(45)
	}
}

func Benchmark_climbStairs1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		climbStairs1(45)
	}
}
