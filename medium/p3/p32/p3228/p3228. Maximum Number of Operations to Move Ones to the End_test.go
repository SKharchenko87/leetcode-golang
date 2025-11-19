package p3228

import "testing"

func Test_maxOperations(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "1001101"}, 4},
		{"Example 2", args{s: "00111"}, 0},
		{"Testcase 345", args{s: "110"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxOperations(tt.args.s); got != tt.want {
				t.Errorf("maxOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxOperations(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxOperations("0011010101100")
	}
}

func Benchmark_maxOperations0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxOperations0("0011010101100")
	}
}
