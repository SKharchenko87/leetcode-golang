package p0904

import "testing"

func Test_totalFruit(t *testing.T) {
	type args struct {
		fruits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{fruits: []int{1, 2, 1}}, 3},
		{"Example 2", args{fruits: []int{0, 1, 2, 2}}, 3},
		{"Example 3", args{fruits: []int{1, 2, 3, 2, 2}}, 4},
		{"TestCase 37", args{fruits: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}}, 5},
		{"TestCase 37", args{fruits: []int{1, 0, 3, 4, 3}}, 3},
		{"TestCase 52", args{fruits: []int{1, 0, 2, 3, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(tt.args.fruits); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_totalFruit(b *testing.B) {
	b.ReportAllocs()
	f := totalFruit
	for i := 0; i < b.N; i++ {
		f([]int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2})
		f([]int{1, 2, 3, 2, 2, 3, 3, 2, 1, 1, 2, 3, 3, 2, 2, 1, 1, 3})
		f([]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7})
		f([]int{3, 3, 3, 1, 2, 1, 2, 1, 2, 3, 1, 1, 2, 2, 1, 3, 3, 2, 2, 3, 1, 2})
		f([]int{4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5})
		f([]int{1, 2, 3, 2, 3, 2, 3, 2, 4, 4, 4, 5, 5, 5, 4, 4, 3, 3, 2, 2})
		f([]int{8, 9, 8, 9, 8, 7, 9, 8, 9, 8, 9, 8, 9, 7, 8, 9, 8, 9, 8})
		f([]int{1, 1, 2, 3, 2, 3, 3, 3, 2, 2, 4, 4, 4, 4, 4, 3, 3, 2, 2, 2})
	}
}

func Benchmark_totalFruit0(b *testing.B) {
	b.ReportAllocs()
	f := totalFruit0
	for i := 0; i < b.N; i++ {
		f([]int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2})
		f([]int{1, 2, 3, 2, 2, 3, 3, 2, 1, 1, 2, 3, 3, 2, 2, 1, 1, 3})
		f([]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7})
		f([]int{3, 3, 3, 1, 2, 1, 2, 1, 2, 3, 1, 1, 2, 2, 1, 3, 3, 2, 2, 3, 1, 2})
		f([]int{4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5})
		f([]int{1, 2, 3, 2, 3, 2, 3, 2, 4, 4, 4, 5, 5, 5, 4, 4, 3, 3, 2, 2})
		f([]int{8, 9, 8, 9, 8, 7, 9, 8, 9, 8, 9, 8, 9, 7, 8, 9, 8, 9, 8})
		f([]int{1, 1, 2, 3, 2, 3, 3, 3, 2, 2, 4, 4, 4, 4, 4, 3, 3, 2, 2, 2})
	}
}
