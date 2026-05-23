package p2833

import "testing"

func Test_furthestDistanceFromOrigin(t *testing.T) {
	type args struct {
		moves string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{moves: "L_RL__R"}, 3},
		{"Example 2", args{moves: "_R__LL_"}, 5},
		{"Example 3", args{moves: "_______"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := furthestDistanceFromOrigin(tt.args.moves); got != tt.want {
				t.Errorf("furthestDistanceFromOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func bench(b *testing.B, f func(s string) int) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f("L_RL__R_R__LL_________R__LL____L_RL__R_____R_R__LL_L_RL__R_R__LL_________R__LL____L_RL__R_____R_R__LL__L_RL__R_____R_R__LL_L_RL__R_R__LL_________R__LL____L_RL__R__")
	}
}

func Benchmark_furthestDistanceFromOrigin(b *testing.B) {
	bench(b, furthestDistanceFromOrigin)
}

func Benchmark_furthestDistanceFromOrigin0(b *testing.B) {
	bench(b, furthestDistanceFromOrigin0)
}
