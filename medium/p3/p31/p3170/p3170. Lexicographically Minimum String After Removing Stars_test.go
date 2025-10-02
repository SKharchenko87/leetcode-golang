package p3170

import "testing"

func Test_clearStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s: "aaba*"}, "aab"},
		{"Example 2", args{s: "abc"}, "abc"},
		{"TestCase 15", args{s: "d*"}, ""},
		{"My 0", args{s: "d"}, "d"},
		{"My 1", args{s: "*"}, ""},
		{"My 2", args{s: "*d"}, "d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clearStars(tt.args.s); got != tt.want {
				t.Errorf("clearStars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_clearStars(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		clearStars("aaba*aslhgfahbank*acndahsdnv*clsadvsa*nvnhasfdgv")
	}
}

func Benchmark_clearStars0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		clearStars0("aaba*aslhgfahbank*acndahsdnv*clsadvsa*nvnhasfdgv")
	}
}
