package p1128

import "testing"

func Test_numEquivDominoPairs(t *testing.T) {
	type args struct {
		dominoes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{dominoes: [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}}, 1},
		{"Example 2", args{dominoes: [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numEquivDominoPairs(tt.args.dominoes); got != tt.want {
				t.Errorf("numEquivDominoPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_numEquivDominoPairs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		numEquivDominoPairs([][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}})
	}
}

func Benchmark_numEquivDominoPairs2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		numEquivDominoPairs2([][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}})
	}
}

func Benchmark_numEquivDominoPairs1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		numEquivDominoPairs1([][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}})
	}
}

func Benchmark_numEquivDominoPairs0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		numEquivDominoPairs0([][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}, {1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}})
	}
}
