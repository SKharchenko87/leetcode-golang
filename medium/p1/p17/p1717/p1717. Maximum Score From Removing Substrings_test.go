package p1717

import "testing"

func Test_maximumGain(t *testing.T) {
	type args struct {
		s string
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "cdbcbbaaabab", x: 4, y: 5}, 19},
		{"Example 2", args{s: "aabbaaxybbaabb", x: 5, y: 4}, 20},
		{"TestCase 7", args{s: "aabbabkbbbfvybssbtaobaaaabataaadabbbmakgabbaoapbbbbobaabvqhbbzbbkapabaavbbeghacabamdpaaqbqabbjbababmbakbaabajabasaabbwabrbbaabbafubayaazbbbaababbaaha", x: 1926, y: 4320}, 112374},
		{"TestCase 61", args{s: "aabbrtababbabmaaaeaabeawmvaataabnaabbaaaybbbaabbabbbjpjaabbtabbxaaavsmmnblbbabaeuasvababjbbabbabbasxbbtgbrbbajeabbbfbarbagha", x: 8484, y: 4096}, 198644},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGain(tt.args.s, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("maximumGain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maximumGain(b *testing.B) {
	b.ReportAllocs()
	f := maximumGain
	for i := 0; i < b.N; i++ {
		f("aabbrtababbabmaaaeaabeawmvaataabnaabbaaaybbbaabbabbbjpjaabbtabbxaaavsmmnblbbabaeuasvababjbbabbabbasxbbtgbrbbajeabbbfbarbagha", 8484, 4096)
		f("cdbcbbaaabab", 4, 5)
		f("aabbabkbbbfvybssbtaobaaaabataaadabbbmakgabbaoapbbbbobaabvqhbbzbbkapabaavbbeghacabamdpaaqbqabbjbababmbakbaabajabasaabbwabrbbaabbafubayaazbbbaababbaaha", 1926, 4320)
		f("ababababababababababababababababababababababababababababababababababababababababababababababababab", 3, 2)
		f("babababababababababababababababababababababababababababababababababababababababababababababababababababa", 6, 7)
		f("aabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabba", 8, 9)
		f("aaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbb", 10, 1)
		f("aabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbab", 9, 11)
		f("bbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabb", 8, 9)
	}
}

func Benchmark_maximumGain1(b *testing.B) {
	b.ReportAllocs()
	f := maximumGain1
	for i := 0; i < b.N; i++ {
		f("aabbrtababbabmaaaeaabeawmvaataabnaabbaaaybbbaabbabbbjpjaabbtabbxaaavsmmnblbbabaeuasvababjbbabbabbasxbbtgbrbbajeabbbfbarbagha", 8484, 4096)
		f("cdbcbbaaabab", 4, 5)
		f("aabbabkbbbfvybssbtaobaaaabataaadabbbmakgabbaoapbbbbobaabvqhbbzbbkapabaavbbeghacabamdpaaqbqabbjbababmbakbaabajabasaabbwabrbbaabbafubayaazbbbaababbaaha", 1926, 4320)
		f("ababababababababababababababababababababababababababababababababababababababababababababababababab", 3, 2)
		f("babababababababababababababababababababababababababababababababababababababababababababababababababababa", 6, 7)
		f("aabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabba", 8, 9)
		f("aaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbb", 10, 1)
		f("aabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbab", 9, 11)
		f("bbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabb", 8, 9)
	}
}

func Benchmark_maximumGain0(b *testing.B) {
	b.ReportAllocs()
	f := maximumGain0
	for i := 0; i < b.N; i++ {
		f("aabbrtababbabmaaaeaabeawmvaataabnaabbaaaybbbaabbabbbjpjaabbtabbxaaavsmmnblbbabaeuasvababjbbabbabbasxbbtgbrbbajeabbbfbarbagha", 8484, 4096)
		f("cdbcbbaaabab", 4, 5)
		f("aabbabkbbbfvybssbtaobaaaabataaadabbbmakgabbaoapbbbbobaabvqhbbzbbkapabaavbbeghacabamdpaaqbqabbjbababmbakbaabajabasaabbwabrbbaabbafubayaazbbbaababbaaha", 1926, 4320)
		f("ababababababababababababababababababababababababababababababababababababababababababababababababab", 3, 2)
		f("babababababababababababababababababababababababababababababababababababababababababababababababababababa", 6, 7)
		f("aabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabba", 8, 9)
		f("aaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbbaaaaaaaaabbbbbbbbbb", 10, 1)
		f("aabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbabbaabbab", 9, 11)
		f("bbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabb", 8, 9)
	}
}
