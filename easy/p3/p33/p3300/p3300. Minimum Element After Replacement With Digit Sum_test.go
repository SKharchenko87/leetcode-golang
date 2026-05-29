package p3300

import "testing"

func Test_minElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{10, 12, 13, 14}}, 1},
		{"Example 2", args{nums: []int{1, 2, 3, 4}}, 1},
		{"Example 3", args{nums: []int{999, 19, 199}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minElement(tt.args.nums); got != tt.want {
				t.Errorf("minElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(b *testing.B, f func([]int) int) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f([]int{91, 143, 215, 335, 412, 458, 642, 680, 711, 792,
			923, 1148, 1205, 1317, 1459, 1582, 1704, 1821, 1965, 2034,
			2110, 2247, 2399, 2415, 2552, 2684, 2713, 2890, 2941, 3056,
			3122, 3289, 3341, 3476, 3510, 3698, 3745, 3812, 3987, 4014,
			4156, 4239, 4371, 4428, 4590, 4613, 4782, 4805, 4934, 5021,
			5144, 5298, 5312, 5467, 5509, 5671, 5723, 5894, 5912, 6045,
			6129, 6274, 6388, 6411, 6595, 6632, 6789, 6814, 6950, 7081,
			7142, 7299, 7351, 7468, 7520, 7694, 7737, 7819, 7980, 8043,
			8167, 8254, 8319, 8472, 8530, 8691, 8724, 8859, 8911, 9053,
			9128, 9246, 9377, 9419, 9580, 9623, 9741, 9805, 9934, 9989,
		})
	}
}

func Benchmark_minElement(b *testing.B) {
	bench(b, minElement)
}

func Benchmark_minElement0(b *testing.B) {
	bench(b, minElement0)
}
