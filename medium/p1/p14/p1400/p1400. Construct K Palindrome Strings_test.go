package p1400

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{s: "annabelle", k: 2}, true},
		{"Example 2", args{s: "leetcode", k: 3}, false},
		{"Example 3", args{s: "true", k: 4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_canConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canConstruct("ejhdyxwbjinykfpvtibzmpjmfjcycunuukipevkuqhmlzeblvbsbpnsvaobtsvsedibnwipzennpchqzrfajmmotyluzleismblxittjepkijxbdmeerlodxzuabftofwzjcmezbrneveikdxsromnljofuaquxdnoxwdfuooqfnflwcmuschzrosvuaollcxocjzlhg", 30)
	}
}

func Benchmark_canConstruct1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canConstruct1("ejhdyxwbjinykfpvtibzmpjmfjcycunuukipevkuqhmlzeblvbsbpnsvaobtsvsedibnwipzennpchqzrfajmmotyluzleismblxittjepkijxbdmeerlodxzuabftofwzjcmezbrneveikdxsromnljofuaquxdnoxwdfuooqfnflwcmuschzrosvuaollcxocjzlhg", 30)
	}
}

func Benchmark_canConstruct0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canConstruct0("ejhdyxwbjinykfpvtibzmpjmfjcycunuukipevkuqhmlzeblvbsbpnsvaobtsvsedibnwipzennpchqzrfajmmotyluzleismblxittjepkijxbdmeerlodxzuabftofwzjcmezbrneveikdxsromnljofuaquxdnoxwdfuooqfnflwcmuschzrosvuaollcxocjzlhg", 30)
	}
}
