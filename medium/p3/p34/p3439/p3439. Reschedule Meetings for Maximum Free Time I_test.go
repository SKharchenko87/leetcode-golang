package p3439

import "testing"

type args struct {
	eventTime int
	k         int
	startTime []int
	endTime   []int
}

func Test_maxFreeTime(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{eventTime: 5, k: 1, startTime: []int{1, 3}, endTime: []int{2, 5}}, 2},
		{"Example 2", args{eventTime: 10, k: 1, startTime: []int{0, 2, 9}, endTime: []int{1, 4, 10}}, 6},
		{"Example 3", args{eventTime: 5, k: 2, startTime: []int{0, 1, 2, 3, 4}, endTime: []int{1, 2, 3, 4, 5}}, 0},
		{"TestCase 350", args{eventTime: 21, k: 1, startTime: []int{7, 10, 16}, endTime: []int{10, 14, 18}}, 7},
		{"TestCase 511", args{eventTime: 21, k: 2, startTime: []int{18, 20}, endTime: []int{20, 21}}, 18},
		{"TestCase 551", args{eventTime: 34, k: 2, startTime: []int{0, 17}, endTime: []int{14, 19}}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFreeTime(tt.args.eventTime, tt.args.k, tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("maxFreeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

var BenchParam args

func init() {
	BenchParam.eventTime = 22
	BenchParam.k = 2
	BenchParam.startTime = []int{3, 7, 9, 13, 21}
	BenchParam.endTime = []int{5, 8, 12, 18, 22}

}

func Benchmark_maxFreeTime(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxFreeTime(BenchParam.eventTime, BenchParam.k, BenchParam.startTime, BenchParam.endTime)
	}
}

func Benchmark_maxFreeTime3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxFreeTime3(BenchParam.eventTime, BenchParam.k, BenchParam.startTime, BenchParam.endTime)
	}
}

func Benchmark_maxFreeTime2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxFreeTime2(BenchParam.eventTime, BenchParam.k, BenchParam.startTime, BenchParam.endTime)
	}
}

func Benchmark_maxFreeTime1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxFreeTime1(BenchParam.eventTime, BenchParam.k, BenchParam.startTime, BenchParam.endTime)
	}
}

func Benchmark_maxFreeTime0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxFreeTime0(BenchParam.eventTime, BenchParam.k, BenchParam.startTime, BenchParam.endTime)
	}
}
