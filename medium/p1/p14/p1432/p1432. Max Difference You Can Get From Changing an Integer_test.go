package p1432

import "testing"

func Test_maxDiff(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{num: 555}, 888},
		{"Example 2", args{num: 9}, 8},
		{"My 1", args{num: 12890}, 82000},
		{"My 2", args{num: 92890}, 87080},
		{"TestCase 207", args{num: 1101057}, 8808050},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDiff(tt.args.num); got != tt.want {
				t.Errorf("maxDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxDiff(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxDiff(1101057)
	}
}

//func Benchmark_maxDiff2(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		maxDiff2(1101057)
//	}
//}

func Benchmark_Approach1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxDiff0(1101057)
	}
}

func Benchmark_Approach2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		maxDiff1(1101057)
	}
}
