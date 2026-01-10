package p0712

import "testing"

func Test_minimumDeleteSum(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s1: "sea", s2: "eat"}, 231},
		{"Example 2", args{s1: "delete", s2: "leet"}, 403},
		{"TestCase 64", args{s1: "a", s2: "at"}, 116},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeleteSum(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(f func(string, string) int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f("nmnmajwmnmnmnmnmnmamsndmasnmdnamnwmndmawnmfwamgnwajkghwaijghanfnasmdnawhoirquoweruqworq", "janskjdhajkdnasndnbjghjgabasdasdashdjaskdhasjdhasjkdhakshdssasdasdasdsa")
	}
}

func Benchmark_minimumDeleteSum1(b *testing.B) {
	bench(minimumDeleteSum1, b)
}

func Benchmark_minimumDeleteSum0(b *testing.B) {
	bench(minimumDeleteSum0, b)
}

func Benchmark_minimumDeleteSum(b *testing.B) {
	bench(minimumDeleteSum, b)
}
