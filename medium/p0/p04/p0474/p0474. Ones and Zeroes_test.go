package p0474

import "testing"

func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{strs: []string{"10", "0001", "111001", "1", "0"}, m: 5, n: 3}, 4},
		{"Example 2", args{strs: []string{"10", "0", "1"}, m: 1, n: 1}, 2},
		{"Test 0", args{strs: []string{"10001111", "111111110001", "10000011001", "111111111", "0000000000"}, m: 5, n: 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findMaxForm(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findMaxForm([]string{"10001111", "111111110001", "10000011001", "111111111", "0000000000"}, 5, 3)
	}
}

func Benchmark_findMaxForm1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findMaxForm1([]string{"10001111", "111111110001", "10000011001", "111111111", "0000000000"}, 5, 3)
	}
}

func Benchmark_findMaxForm0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findMaxForm0([]string{"10001111", "111111110001", "10000011001", "111111111", "0000000000"}, 5, 3)
	}
}
