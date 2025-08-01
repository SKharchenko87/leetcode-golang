package p2044

import "testing"

func Test_countMaxOrSubsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{nums: []int{3, 1}}, 2},
		{"Example 2", args{nums: []int{2, 2, 2}}, 7},
		{"Example 3", args{nums: []int{3, 2, 1, 5}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMaxOrSubsets(tt.args.nums); got != tt.want {
				t.Errorf("countMaxOrSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countMaxOrSubsets(b *testing.B) {
	b.ReportAllocs()
	var f func([]int) int
	f = countMaxOrSubsets
	for i := 0; i < b.N; i++ {
		f([]int{12007})
		f([]int{39317, 28735, 71230, 59313, 19954})
		f([]int{35569, 91997, 54930, 66672, 12363})
		f([]int{32, 39, 37, 31, 42, 38, 29, 43, 40, 36, 44, 35, 41, 33, 34, 30})
		f([]int{63609, 94085, 69432, 15248, 22060, 76843, 84075, 835, 23463, 66399, 95031, 22676, 92115})
		f([]int{6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890})
		f([]int{90193, 56697, 77229, 72927, 74728, 93652, 70751, 32415, 94774, 9067, 14758, 59835, 26047, 36393, 60530, 64649})
		f([]int{89260, 58129, 16949, 64128, 9782, 26664, 96968, 59838, 27627, 68561, 4026, 91345, 26966, 28876, 46276, 19878})
	}
}

func Benchmark_countMaxOrSubsets3(b *testing.B) {
	b.ReportAllocs()
	var f func([]int) int
	f = countMaxOrSubsets3
	for i := 0; i < b.N; i++ {
		f([]int{12007})
		f([]int{39317, 28735, 71230, 59313, 19954})
		f([]int{35569, 91997, 54930, 66672, 12363})
		f([]int{32, 39, 37, 31, 42, 38, 29, 43, 40, 36, 44, 35, 41, 33, 34, 30})
		f([]int{63609, 94085, 69432, 15248, 22060, 76843, 84075, 835, 23463, 66399, 95031, 22676, 92115})
		f([]int{6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890})
		f([]int{90193, 56697, 77229, 72927, 74728, 93652, 70751, 32415, 94774, 9067, 14758, 59835, 26047, 36393, 60530, 64649})
		f([]int{89260, 58129, 16949, 64128, 9782, 26664, 96968, 59838, 27627, 68561, 4026, 91345, 26966, 28876, 46276, 19878})
	}
}

func Benchmark_countMaxOrSubsets2(b *testing.B) {
	b.ReportAllocs()
	var f func([]int) int
	f = countMaxOrSubsets2
	for i := 0; i < b.N; i++ {
		f([]int{12007})
		f([]int{39317, 28735, 71230, 59313, 19954})
		f([]int{35569, 91997, 54930, 66672, 12363})
		f([]int{32, 39, 37, 31, 42, 38, 29, 43, 40, 36, 44, 35, 41, 33, 34, 30})
		f([]int{63609, 94085, 69432, 15248, 22060, 76843, 84075, 835, 23463, 66399, 95031, 22676, 92115})
		f([]int{6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890})
		f([]int{90193, 56697, 77229, 72927, 74728, 93652, 70751, 32415, 94774, 9067, 14758, 59835, 26047, 36393, 60530, 64649})
		f([]int{89260, 58129, 16949, 64128, 9782, 26664, 96968, 59838, 27627, 68561, 4026, 91345, 26966, 28876, 46276, 19878})
	}
}

func Benchmark_countMaxOrSubsets1(b *testing.B) {
	b.ReportAllocs()
	var f func([]int) int
	f = countMaxOrSubsets1
	for i := 0; i < b.N; i++ {
		f([]int{12007})
		f([]int{39317, 28735, 71230, 59313, 19954})
		f([]int{35569, 91997, 54930, 66672, 12363})
		f([]int{32, 39, 37, 31, 42, 38, 29, 43, 40, 36, 44, 35, 41, 33, 34, 30})
		f([]int{63609, 94085, 69432, 15248, 22060, 76843, 84075, 835, 23463, 66399, 95031, 22676, 92115})
		f([]int{6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890})
		f([]int{90193, 56697, 77229, 72927, 74728, 93652, 70751, 32415, 94774, 9067, 14758, 59835, 26047, 36393, 60530, 64649})
		f([]int{89260, 58129, 16949, 64128, 9782, 26664, 96968, 59838, 27627, 68561, 4026, 91345, 26966, 28876, 46276, 19878})
	}
}

func Benchmark_countMaxOrSubsets0(b *testing.B) {
	b.ReportAllocs()
	var f func([]int) int
	f = countMaxOrSubsets0
	for i := 0; i < b.N; i++ {
		f([]int{12007})
		f([]int{39317, 28735, 71230, 59313, 19954})
		f([]int{35569, 91997, 54930, 66672, 12363})
		f([]int{32, 39, 37, 31, 42, 38, 29, 43, 40, 36, 44, 35, 41, 33, 34, 30})
		f([]int{63609, 94085, 69432, 15248, 22060, 76843, 84075, 835, 23463, 66399, 95031, 22676, 92115})
		f([]int{6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890, 6890})
		f([]int{90193, 56697, 77229, 72927, 74728, 93652, 70751, 32415, 94774, 9067, 14758, 59835, 26047, 36393, 60530, 64649})
		f([]int{89260, 58129, 16949, 64128, 9782, 26664, 96968, 59838, 27627, 68561, 4026, 91345, 26966, 28876, 46276, 19878})
	}
}
