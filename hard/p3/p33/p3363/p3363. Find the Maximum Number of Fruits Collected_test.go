package p3363

import "testing"

func Test_maxCollectedFruits(t *testing.T) {
	type args struct {
		fruits [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{fruits: [][]int{{1, 2, 3, 4}, {5, 6, 8, 7}, {9, 10, 11, 12}, {13, 14, 15, 16}}}, 100},
		{"Example 2", args{fruits: [][]int{{1, 1}, {1, 1}}}, 4},
		{"My 1", args{fruits: [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}}, 13},
		{"My 2", args{fruits: [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}}, 10},
		{"TestCase 2", args{fruits: [][]int{{8, 5, 0, 17, 15}, {6, 0, 15, 6, 0}, {7, 19, 16, 8, 18}, {11, 3, 2, 12, 13}, {17, 15, 15, 4, 6}}}, 145},
		{"TestCase 283", args{fruits: [][]int{{16, 3, 11, 14, 14}, {3, 0, 10, 13, 14}, {7, 18, 8, 7, 18}, {7, 8, 5, 7, 5}, {0, 14, 8, 1, 0}}}, 105},
		{"TestCase 450", args{fruits: [][]int{{4, 2, 1, 15, 20, 6, 1, 3, 11}, {1, 13, 19, 18, 15, 13, 6, 14, 13}, {20, 12, 10, 8, 7, 12, 20, 5, 17}, {18, 9, 15, 9, 16, 6, 13, 1, 18}, {12, 19, 12, 7, 1, 15, 16, 13, 9}, {17, 19, 15, 17, 16, 11, 14, 4, 15}, {10, 16, 7, 4, 8, 19, 13, 2, 10}, {1, 2, 14, 6, 0, 13, 11, 5, 16}, {16, 17, 12, 8, 2, 12, 19, 4, 7}}}, 288},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCollectedFruits4(tt.args.fruits); got != tt.want {
				t.Errorf("maxCollectedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxCollectedFruits(b *testing.B) {
	b.ReportAllocs()
	f := maxCollectedFruits
	for i := 0; i < b.N; i++ {
		f([][]int{{1, 2, 3, 4}, {5, 6, 8, 7}, {9, 10, 11, 12}, {13, 14, 15, 16}})
		f([][]int{{1, 1}, {1, 1}})
		f([][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}})
		f([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}})
	}
}

func Benchmark_maxCollectedFruits0(b *testing.B) {
	b.ReportAllocs()
	f := maxCollectedFruits0
	for i := 0; i < b.N; i++ {
		f([][]int{{1, 2, 3, 4}, {5, 6, 8, 7}, {9, 10, 11, 12}, {13, 14, 15, 16}})
		f([][]int{{1, 1}, {1, 1}})
		f([][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}})
		f([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}})
	}
}

func Benchmark_maxCollectedFruits1(b *testing.B) {
	b.ReportAllocs()
	f := maxCollectedFruits1
	for i := 0; i < b.N; i++ {
		f([][]int{{1, 2, 3, 4}, {5, 6, 8, 7}, {9, 10, 11, 12}, {13, 14, 15, 16}})
		f([][]int{{1, 1}, {1, 1}})
		f([][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}})
		f([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}})
	}
}

func Benchmark_maxCollectedFruits2(b *testing.B) {
	b.ReportAllocs()
	f := maxCollectedFruits2
	for i := 0; i < b.N; i++ {
		f([][]int{{1, 2, 3, 4}, {5, 6, 8, 7}, {9, 10, 11, 12}, {13, 14, 15, 16}})
		f([][]int{{1, 1}, {1, 1}})
		f([][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}})
		f([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}})
	}
}

func Benchmark_maxCollectedFruits3(b *testing.B) {
	b.ReportAllocs()
	f := maxCollectedFruits3
	for i := 0; i < b.N; i++ {
		f([][]int{{1, 2, 3, 4}, {5, 6, 8, 7}, {9, 10, 11, 12}, {13, 14, 15, 16}})
		f([][]int{{1, 1}, {1, 1}})
		f([][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}})
		f([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}})
	}
}

func Benchmark_maxCollectedFruits4(b *testing.B) {
	b.ReportAllocs()
	f := maxCollectedFruits4
	for i := 0; i < b.N; i++ {
		f([][]int{{1, 2, 3, 4}, {5, 6, 8, 7}, {9, 10, 11, 12}, {13, 14, 15, 16}})
		f([][]int{{1, 1}, {1, 1}})
		f([][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}})
		f([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}})
	}
}
