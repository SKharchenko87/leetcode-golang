package p0326

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{n: 27}, true},
		{"Example 2", args{n: 0}, false},
		{"Example 3", args{n: -1}, false},
		{"TestCase 21039", args{n: 1162261467}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isPowerOfThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPowerOfThree(27)
		isPowerOfThree(1162261467)
	}
}

func Benchmark_isPowerOfThree1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPowerOfThree1(27)
		isPowerOfThree1(1162261467)
	}
}

func Benchmark_isPowerOfThree0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPowerOfThree0(27)
		isPowerOfThree0(1162261467)
	}
}
