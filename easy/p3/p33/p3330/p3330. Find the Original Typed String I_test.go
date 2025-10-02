package p3330

import "testing"

func Test_possibleStringCount(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{word: "abbcccc"}, 5},
		{"Example 2", args{word: "abcd"}, 1},
		{"Example 3", args{word: "aaaa"}, 4},
		{"Discussion 0", args{word: "xxulxumxtdpkoevrpqzqkiynoopculvmjlfqsexnobtylogwmohqnvnengkqupxscfoygdygvbopnanxickqfggjiqxjanvmcbdb"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleStringCount(tt.args.word); got != tt.want {
				t.Errorf("possibleStringCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_possibleStringCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		possibleStringCount("xxulxumxtdpkoevrpqzqkiynoopculvmjlfqsexnobtylogwmohqnvnengkqupxscfoygdygvbopnanxickqfggjiqxjanvmcbdb")
	}
}

func Benchmark_possibleStringCount0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		possibleStringCount0("xxulxumxtdpkoevrpqzqkiynoopculvmjlfqsexnobtylogwmohqnvnengkqupxscfoygdygvbopnanxickqfggjiqxjanvmcbdb")
	}
}
