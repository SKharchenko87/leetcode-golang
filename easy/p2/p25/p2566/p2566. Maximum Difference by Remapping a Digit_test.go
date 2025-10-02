package p2566

import "testing"

func Test_minMaxDifference(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{num: 11891}, 99009},
		{"Example 2", args{num: 90}, 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMaxDifference(tt.args.num); got != tt.want {
				t.Errorf("minMaxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_minMaxDifference(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		minMaxDifference(11891)
	}
}

func Benchmark_minMaxDifference1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		minMaxDifference1(11891)
	}
}
